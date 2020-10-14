package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

type register struct {
	win  int
	draw int
	loss int
}

func (reg *register) getMatchPlayed() int {
	return reg.win + reg.draw + reg.loss
}

func (reg *register) getPoints() int {
	return 3*reg.win + reg.draw
}

func getPunctuationsTable(reader io.Reader) (error, map[string]*register) {
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
			return err, nil
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

	return nil, m
}

func Tally(reader io.Reader, buffer io.Writer) error {
	err, table := getPunctuationsTable(reader)

	if err != nil {
		return errors.New("Problem with table")
	}

	buffer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))

	for k, v := range table {
		s := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", k, v.getMatchPlayed(), v.win, v.draw, v.loss, v.getPoints())
		buffer.Write([]byte(s))
	}

	return nil
}
