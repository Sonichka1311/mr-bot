package helpers

import "mr-bot/pkg/constants"

var missedQueue []string

func GetNext() string {
	constants.LastPersonNum = (constants.LastPersonNum + 1) % len(constants.People)
	return constants.People[constants.LastPersonNum]
}

// GetDuty returns telegram username of new responsible
func GetDuty(author string) string {
	if len(missedQueue) != 0 {
		for idx, person := range missedQueue {
			if person != author {
				if idx + 1 == len(missedQueue) {
					missedQueue = missedQueue[:idx]
				} else {
					missedQueue = append(missedQueue[:idx], missedQueue[idx + 1:]...)
				}
				return person
			}
		}
	}
	person := GetNext()
	if person != author {
		return person
	}
	missedQueue = append(missedQueue, person)
	return GetNext()
}
