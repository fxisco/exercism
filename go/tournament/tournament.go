package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type register struct {
	win  int
	draw int
	loss int
}

func Tally(reader io.Reader, buffer io.Writer) error {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	m := make(map[string]*register)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		teamA, prsA := m[record[0]]
		teamB, prsB := m[record[1]]

		action := strings.TrimSuffix(record[2], "\n")

		if !prsA {
			teamA = &register{}
		}

		if !prsB {
			teamB = &register{}
		}

		if action == "win" {
			teamA.win++
			teamB.loss++
		} else if action == "loss" {
			teamA.loss++
			teamB.win++
		} else {
			teamA.draw++
			teamB.draw++
		}

		m[record[0]] = teamA
		m[record[1]] = teamB
	}

	for k, v := range m {
		fmt.Printf("key[%s] value[%v]\n", k, v)
	}

	buffer.Write([]byte("Team                           | MP |  W |  D |  L |  P"))

	return nil
}
