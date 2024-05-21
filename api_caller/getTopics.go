package api_caller

// import (
// 	"net/url"
// 	"strconv"
// )

// func GetTopicsChan(pageSize int) (chan string, error) {

// 	topicIDs := make(chan string, 400)

// 	var bodyParams map[string][]byte = map[string][]byte{
// 		"query":     GetTopicQuery(),
// 		"languages": []byte(`["en"]`),
// 	}

// 	var urlParams url.Values = url.Values{
// 		"apiKey":   []string{"SEDIA"},
// 		"text":     []string{"***"},
// 		"pageSize": []string{strconv.Itoa(pageSize)},
// 	}

// 	api_caller, err := NewAPI_Caller(bodyParams, urlParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resultsChan, err := api_caller.GetResults()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// topicsMap := make(map[string]int)
// 	// counter := 0

// 	// file, err := os.Create("topics2.txt")
// 	// if err != nil {
// 	// 	os.Exit(1)
// 	// }

// 	//go func() {

// 	for resultsArr := range resultsChan {
// 		for inx := range resultsArr {

// 			// if _, ok := resultsArr[inx].Metadata["identifier"]; !ok ||
// 			// 	len(resultsArr[inx].Metadata["identifier"]) != 1 {
// 			// 	log.Println("no identifier")
// 			// 	continue
// 			// }

// 			topicIDs <- resultsArr[inx].Metadata["identifier"][0]

// 			// if topicsMap[resultsArr[inx].Metadata["identifier"][0]] != 0 {
// 			// 	fmt.Println(resultsArr[inx].Metadata["identifier"][0])
// 			// 	counter++
// 			// }

// 			// topicsMap[resultsArr[inx].Metadata["identifier"][0]]++

// 			// file.WriteString(resultsArr[inx].Metadata["identifier"][0] + "\n")
// 		}
// 	}

// 	// for id, encounter := range topicsMap {
// 	// 	if encounter != 1 {
// 	// 		fmt.Println(id, encounter)
// 	// 		counter++
// 	// 	}
// 	// }

// 	//fmt.Println(counter)

// 	close(topicIDs)
// 	//}()

// 	return topicIDs, nil
// }
