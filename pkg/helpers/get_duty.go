package helpers

import "mr-bot/pkg/constants"

func GetDuty() string {
	constants.LastPersonNum = (constants.LastPersonNum + 1) % len(constants.People)
	return constants.People[constants.LastPersonNum]
}
