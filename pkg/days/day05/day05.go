package days

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"aoc2024/pkg/common"
)

func Process(content *[]byte) (string, string, error) {
	result1, err := part1(content)
	if err != nil {
		return "", "", err
	}
	result2, err := part2(content)
	if err != nil {
		return result1, "", err
	}
	return result1, result2, nil
}

func part1(content *[]byte) (string, error) {
	lines := common.ReadByLine(content)
	ruleReadingDone := false
	ruleStorage := make(map[int]printJob)
	var updates []printerUpdate
	for _, line := range lines {
		if len(line) == 0 {
			ruleReadingDone = true
			continue
		}
		if !ruleReadingDone {
			upsertPrinterJob(line, ruleStorage)
		} else {
			updates = append(updates, createPrinterUpdate(line))
		}

	}

	sum := 0
	for _, update := range updates {
		// remember, that update is a copy with only scope of for loop!
		// as long as I don't need the values outside, I can keep it like that.
		// otherwise switch to index-access.
		update.runValidation(ruleStorage)
		if update.valid {
			// assuming all have odd length
			sum += update.updateInformation[update.length/2]
		}
	}

	return strconv.Itoa(sum), nil
}

func part2(content *[]byte) (string, error) {
	lines := common.ReadByLine(content)
	ruleReadingDone := false
	ruleStorage := make(map[int]printJob)
	var updates []printerUpdate
	for _, line := range lines {
		if len(line) == 0 {
			ruleReadingDone = true
			continue
		}
		if !ruleReadingDone {
			upsertPrinterJob(line, ruleStorage)
		} else {
			updates = append(updates, createPrinterUpdate(line))
		}

	}

	sum := 0
	for _, update := range updates {
		// remember, that update is a copy with only scope of for loop!
		// as long as I don't need the values outside, I can keep it like that.
		// otherwise switch to index-access.
		update.runValidation(ruleStorage)
		if !update.valid {
			update.orderJobByRuleset(ruleStorage)
			// assuming all have odd length
			sum += update.updateInformation[update.length/2]
		}
	}

	return strconv.Itoa(sum), nil
}

type printJob struct {
	id    int
	needs []int
}

func upsertPrinterJob(s string, storage map[int]printJob) {
	splits := strings.Split(s, "|")
	size := len(splits)
	if size > 2 {
		msg := fmt.Sprintf("unexpected amount of splits. there should be only two but there are %d. (original='%s')", len(splits), s)
		slog.Error(msg)
		panic(msg)
	}
	needed, err := strconv.Atoi(splits[0])
	if err != nil {
		msg := fmt.Sprintf("error on transforming to integer (value='%s')", splits[0])
		slog.Error(msg, "error", err)
		panic(msg)
	}
	id, err := strconv.Atoi(splits[1])
	if err != nil {
		msg := fmt.Sprintf("error on transforming to integer (value='%s')", splits[1])
		slog.Error(msg, "error", err)
		panic(msg)
	}

	foundJob, exists := storage[id]

	if !exists {
		job := printJob{
			id:    id,
			needs: []int{needed},
		}
		storage[id] = job
	} else {
		tmp := append(foundJob.needs, needed)
		foundJob.needs = tmp
		storage[id] = foundJob
	}
}

type printerUpdate struct {
	length            int
	currentPosition   int
	currentValue      int
	updateInformation []int
	valid             bool
}

func (pu *printerUpdate) read() bool {
	if pu.currentPosition+1 < pu.length {
		pu.currentPosition += 1
		pu.currentValue = pu.updateInformation[pu.currentPosition]
		return true
	}
	return false
}

func createPrinterUpdate(s string) printerUpdate {
	splits := strings.Split(s, ",")
	size := len(splits)
	updateInfo := make([]int, size)
	for i, val := range splits {
		var err error
		updateInfo[i], err = strconv.Atoi(val)
		if err != nil {
			slog.Error(fmt.Sprintf("There was an issue converting %s to integer", val), "error", err)
			panic(err)
		}
	}
	return printerUpdate{
		length:            size,
		currentPosition:   -1,
		currentValue:      -1,
		updateInformation: updateInfo,
	}
}

func (pu *printerUpdate) runValidation(ruleStorage map[int]printJob) {
	rulesForUpdate := pu.onlyRulesApplying(ruleStorage)
	seen := make(map[int]int)
	valid := true

	for pu.read() && valid {
		jobId := pu.currentValue
		seen[jobId] = jobId
		job, found := rulesForUpdate[jobId]
		if !found {
			continue
		}
		jobNeeds := job.needs
		for _, neededJobId := range jobNeeds {
			if !slices.Contains(jobNeeds, neededJobId) {
				continue
			}
			_, wasSeen := seen[neededJobId]
			valid = valid && wasSeen
		}
	}
	pu.valid = valid
	pu.resetRead()
}

func (pu *printerUpdate) orderJobByRuleset(ruleStorage map[int]printJob) {
	rulesForUpdate := pu.onlyRulesApplying(ruleStorage)

	ids := make([]int, len(pu.updateInformation))
	copy(ids, pu.updateInformation)

	slices.SortStableFunc(ids, func(a, b int) int {
		jobA, found := rulesForUpdate[a]
		sizeNeedsA := 0
		if found {
			sizeNeedsA = len(jobA.needs)
		}

		jobB, found := rulesForUpdate[b]
		sizeNeedsB := 0
		if found {
			sizeNeedsB = len(jobB.needs)
		}

		return sizeNeedsA - sizeNeedsB
	})
	copy(pu.updateInformation, ids)
}

func (pu *printerUpdate) resetRead() {
	pu.currentValue = -1
	pu.currentPosition = -1
}

func (pu *printerUpdate) onlyRulesApplying(ruleStorage map[int]printJob) map[int]printJob {
	rulesThatApply := make(map[int]printJob)
	for _, idInUpdate := range pu.updateInformation {
		job, found := ruleStorage[idInUpdate]
		if !found {
			continue
		}
		newJob := printJob{id: job.id}
		for _, id := range job.needs {
			contains := slices.Contains(pu.updateInformation, id)
			if contains {
				newJob.needs = append(newJob.needs, id)
			}
		}
		rulesThatApply[idInUpdate] = newJob
	}
	return rulesThatApply
}
