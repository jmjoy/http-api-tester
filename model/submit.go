package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bytes"

	"github.com/jmjoy/boomer"
)

var SubmitModel = &submitModel{}

type submitModel struct{}

func (this *submitModel) Submit(data Data) (resp Response, err error) {
	if err = data.Validate(); err != nil {
		return
	}

	data, err = HookPlugin(data)
	if err != nil {
		return
	}

	reqMaker, err := NewRequestMaker(data)
	if err != nil {
		return
	}

	var response Response

	if err = this.submitTest(data, reqMaker, &response); err != nil {
		return
	}

	if err = this.submitBenckmark(data, data.Bm, reqMaker, &response); err != nil {
		return
	}

	// save to history
	if err = HistoryModel.Insert(data); err != nil {
		return
	}

	return response, nil
}

func (this *submitModel) submitTest(data Data, reqMaker *RequestMaker, response *Response) error {
	req, err := reqMaker.NewRequest()
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response.ReqUrl = req.URL.String()
	response.ReqBody = reqMaker.Body
	response.Status = resp.Status
	response.Test = string(respBodyBuf)

	return nil
}

func (this *submitModel) submitBenckmark(data Data, bm Bm, reqMaker *RequestMaker, response *Response) error {
	if !bm.Switch {
		return nil
	}

	// limit N, C reduce server pressure.
	var n, c = bm.N, bm.C
	if bm.N >= 1000 {
		n = 1000
	}
	if bm.C >= 500 {
		c = 500
	}

	req, err := reqMaker.NewRequest()
	if err != nil {
		return err
	}

	result := (&boomer.Boomer{
		Request:     req,
		RequestBody: reqMaker.Body,
		N:           int(n),
		C:           int(c),
		Timeout:     35,
	}).Run()

	response.Bm = formatReportResult(result)

	return nil
}

func formatReportResult(result *boomer.ReportResult) string {
	buffer := new(bytes.Buffer)

	buffer.WriteString("\nSummary:\n")
	buffer.WriteString(fmt.Sprintf("  Total:\t%4.4f secs.\n", result.Summary.TotalSecond))
	buffer.WriteString(fmt.Sprintf("  Slowest:\t%4.4f secs.\n", result.Summary.SlowestSecond))
	buffer.WriteString(fmt.Sprintf("  Fastest:\t%4.4f secs.\n", result.Summary.FastestSecond))
	buffer.WriteString(fmt.Sprintf("  Average:\t%4.4f secs.\n", result.Summary.AverageSecond))
	buffer.WriteString(fmt.Sprintf("  Requests/sec:\t%4.4f\n", result.Summary.RequestsPerSec))
	if result.Summary.TotalSize > 0 {
		buffer.WriteString(fmt.Sprintf("  Total Data Received:\t%d bytes.\n", result.Summary.TotalSize))
		buffer.WriteString(fmt.Sprintf("  Response Size per Request:\t%d bytes.\n", result.Summary.SizePerRequest))
	}

	buffer.WriteString("\nStatus code distribution:\n")
	for code, num := range result.StatusCodeDist {
		buffer.WriteString(fmt.Sprintf("  [%d]\t%d responses\n", code, num))
	}

	buffer.WriteString("\nResponse time histogram:\n")
	for _, v := range result.ResponseTimes {
		buffer.WriteString(fmt.Sprintf("  %4.3f [%v]\t|%v\n", v.Second, v.Count, strings.Repeat("*", v.BarLen)))
	}

	buffer.WriteString("\nLatency distribution:\n")
	for k, v := range result.LatencyDist {
		buffer.WriteString(fmt.Sprintf("  %v%% in %4.4f secs.\n", k, v))
	}

	if len(result.ErrorDist) > 0 {
		buffer.WriteString("\nError distribution:\n")
		for err, num := range result.ErrorDist {
			buffer.WriteString(fmt.Sprintf("  [%d]\t%s\n", num, err))
		}
	}

	return buffer.String()
}
