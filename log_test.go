package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	traceText   = "trace text"
	debugText   = "debug text"
	infoText    = "info text"
	warningText = "warning text"
	errorText   = "error text"
	fatalText   = "fatal text"
)

func TestInitCommonMode(t *testing.T) {
	nameFile := "commonLog.log"
	defer os.Remove(nameFile)

	if err := Init(nameFile, CommonLog, false); err != nil {
		t.Errorf("Init() error = %v", err)
	}

	Trace.Println(traceText)
	Debug.Println(debugText)
	Info.Println(infoText)
	Warning.Println(warningText)
	Error.Println(errorText)
	Fatal.Println(fatalText)

	expectedLogs := []string{
		infoText,
		warningText,
		errorText,
		fatalText,
	}

	unexpectedLogs := []string{
		traceText,
		debugText,
	}

	file, err := os.Open(nameFile)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
	data := string(dataBytes)

	for _, expectedLog := range expectedLogs {
		if !strings.Contains(data, expectedLog) {
			t.Errorf("Incorrect logs, expected to find: %v in  \n%v", expectedLog, data)
		}
	}

	for _, unexpectedLog := range unexpectedLogs {
		if strings.Contains(data, unexpectedLog) {
			t.Errorf("Incorrect logs, unexpected to find: %v in \n%v", unexpectedLog, data)
		}
	}
}

func TestInitTraceMode(t *testing.T) {
	nameFile := "traceLog.log"
	defer os.Remove(nameFile)

	if err := Init(nameFile, TraceLog, false); err != nil {
		t.Errorf("Init() error = %v", err)
	}

	Trace.Println(traceText)
	Debug.Println(debugText)
	Info.Println(infoText)
	Warning.Println(warningText)
	Error.Println(errorText)
	Fatal.Println(fatalText)

	expectedLogs := []string{
		traceText,
		debugText,
		infoText,
		warningText,
		errorText,
		fatalText,
	}

	file, err := os.Open(nameFile)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
	data := string(dataBytes)

	for _, expectedLog := range expectedLogs {
		if !strings.Contains(data, expectedLog) {
			t.Errorf("Incorrect logs, expected to find: %v in  \n%v", expectedLog, data)
		}
	}
}

func TestInitTestLogs(t *testing.T) {
	nameFile := "commonTestLog.log"
	defer os.Remove(nameFile)

	if err := Init(nameFile, CommonLog, true); err != nil {
		t.Errorf("Init() error = %v", err)
	}

	Trace.Println(traceText)
	Debug.Println(debugText)
	Info.Println(infoText)
	Warning.Println(warningText)
	Error.Println(errorText)
	Fatal.Println(fatalText)

	file, err := os.Open(nameFile)
	if err == nil {
		file.Close()
		t.Error("File was not expected to exist")
	}
}
