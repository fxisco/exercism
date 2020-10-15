package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type register struct {
	name string
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

type ByPoints []*register

func (a ByPoints) Len() int      { return len(a) }
func (a ByPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool {
	switch {
	case a[i].getPoints() != a[j].getPoints():
		return a[i].getPoints() > a[j].getPoints()
	case a[i].win != a[j].win:
		return a[i].win > a[j].win
	default:
		return a[i].name < a[j].name
	}
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

		if len(record) != 3 {
			return errors.New("INPUT: Bad format"), nil
		}

		teamA, prsA := m[record[0]]
		teamB, prsB := m[record[1]]

		action := strings.TrimSuffix(record[2], "\n")

		if !prsA {
			teamA = &register{name: record[0]}
		}

		if !prsB {
			teamB = &register{name: record[1]}
		}

		if action == "win" {
			teamA.win++
			teamB.loss++
		} else if action == "loss" {
			teamA.loss++
			teamB.win++
		} else if action == "draw" {
			teamA.draw++
			teamB.draw++
		} else {
			return errors.New("GAME RESULT: Bad format"), nil
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

	// convert map to array
	values := []*register{}

	for _, value := range table {
		values = append(values, value)
	}

	sort.Sort(ByPoints(values))

	for _, v := range values {
		s := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", v.name, v.getMatchPlayed(), v.win, v.draw, v.loss, v.getPoints())
		buffer.Write([]byte(s))
	}

	return nil
}
