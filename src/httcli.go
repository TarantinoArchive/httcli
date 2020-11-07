package main

import (
	"io"
	"os"
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"github.com/rivo/tview"
)


// Structs of requests to add in the slice requestHistory
type headerStruct struct {
	name string
	content string
}
type requestStruct struct {
	host string
	bodyPath string
	requestType string
	headers []headerStruct
}


// Function check if path exists and is not a directory
func checkFileValidity(path string) int {
	fileinfo, err := os.Stat(path)
    if os.IsNotExist(err) || fileinfo.IsDir() {
        return 1
	} else if err != nil {
		return 2
	}
	return 0
}

// Function to get the index of an element from a slice
func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
 }

func main() {


	var requestHistory []requestStruct
	requestTypes := []string{"GET", "PUT", "POST", "PATCH", "DELETE"}

	// Creating an empty Text View to fill later with errors or the response text
	responseText := tview.NewTextView()
	responseText.SetDynamicColors(true)

	// Creating a Dropdown to select the Request Type between requestTypes
	requestSelector := tview.NewDropDown().
						SetLabel("Request method ").
						SetOptions(requestTypes, nil)

	// Creating a Text Field to get the file path of the file to use as body, and the host
	bodyFilePath := tview.NewInputField().SetLabel("File path ")
	host := tview.NewInputField().SetLabel("Host ")

	// Creating an empty header list and a form to add headers to header list
	headersList := tview.NewList()
	addHeader := tview.NewForm().
		AddInputField("Name ", "", 20, nil, nil).
		AddInputField("Content ", "", 20, nil, nil)
	addHeader.AddButton("Add to headers", func() {
			nameFieldText := addHeader.GetFormItemByLabel("Name ").(*tview.InputField).GetText()
			contentFieldText := addHeader.GetFormItemByLabel("Content ").(*tview.InputField).GetText()
			headersList.AddItem(nameFieldText, contentFieldText, ' ', func() {
				headersList.RemoveItem(headersList.GetCurrentItem())
			})
	})

	// Creating an empty history list. Every time the user make a request, the request will be added to historyList
	historyList := tview.NewList()

	sendRequestButton := tview.NewButton("Send Request").SetSelectedFunc(func() {

		// Getting the values of all the text fields
		_, requestTypeText := requestSelector.GetCurrentOption()
		hostText := host.GetText()
		bodyPathText := bodyFilePath.GetText()

		// Getting the content of the file. Body empty if no file or file invalid.
		// Convert body to a Reader because the body field of the NewRequest function requires a Reader
		var bodyText io.Reader
		if checkFileValidity(bodyPathText)!=0 {
			bodyText = nil
		} else {
			content, err := ioutil.ReadFile(bodyPathText)
			if err != nil {
				responseText.SetText("[#ff0000]There was an error reading the body file. \nError log below.\n\n"+err.Error()+"[#ff0000]")
				return
			}
			bodyText = bytes.NewReader(content)
		}

		// Creating request variable, to add to the history of the requests
		request := requestStruct{hostText, bodyPathText, requestTypeText, make([]headerStruct, 0)}

		// Composing the request
		client := &http.Client{}
		req, err := http.NewRequest(requestTypeText, hostText, bodyText)
		if err != nil {
			responseText.SetText("[#ff0000]There was an error composing the request. \nError log below.\n\n"+err.Error()+"[#ff0000]")
			return
		}
		// Adding the headers to the request
		for i:=0; i < headersList.GetItemCount(); i++ {
			headerName, headerContent := headersList.GetItemText(i)
			req.Header.Add(headerName, headerContent)
			request.headers = append(request.headers, headerStruct{headerName, headerContent})
		}

		// Sending the request
		resp, err := client.Do(req)
		if err != nil {
			responseText.SetText("[#ff0000]There was an error doing the request. \nError log below.\n\n"+err.Error()+"[#ff0000]")
			return
		}

		// Composing the output to put
		viewText := "Status: " + resp.Status + "\n\nHeaders:\n"
		for k, header := range resp.Header {
			viewText = viewText + k + ": " + header[0] + "\n"
		}
		// Converting the body of type Reader to Bytes, and then string to add the string to viewText
		responseBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			responseText.SetText("[#ff0000]There was an error getting the response body. \nError log below.\n\n"+err.Error()+"[#ff0000]")
			return
		}
		viewText = viewText + "\nBody: \n" + string(responseBytes[:])

		// Adding to historyList and setting the function to set every field to the selected item in the history
		historyList.AddItem(hostText, requestTypeText, []rune(strconv.Itoa(len(requestHistory)))[0], func() {
			headersList.Clear()
			index := historyList.GetCurrentItem()
			host.SetText(requestHistory[index].host)
			bodyFilePath.SetText(requestHistory[index].bodyPath)
			requestSelector.SetCurrentOption(indexOf(requestHistory[index].requestType, requestTypes))
			for _, header := range requestHistory[index].headers {
				headersList.AddItem(header.name, header.content, ' ', nil)
			}
		})
		requestHistory = append(requestHistory, request) 

		responseText.SetText(viewText)
	})

	sendRequestButton.SetBorder(true).SetRect(0, 0, 3, 1)
	grid := tview.NewGrid().
		SetRows(3, 3, 3, 3, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(requestSelector, 0, 0, 1, 1, 0, 100, false).
		AddItem(host, 1, 0, 1, 1, 0, 100, false).
		AddItem(bodyFilePath, 2, 0, 1, 1, 0, 100, false).
		AddItem(responseText, 0, 1, 5, 1, 0, 100, false).
		AddItem(sendRequestButton, 4, 0, 1, 1, 0, 100, false).
		AddItem(historyList, 3, 0, 1, 1, 0, 0, false).
		AddItem(headersList, 0, 2, 3, 1, 0, 0, false).
		AddItem(addHeader, 3, 2, 2, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}