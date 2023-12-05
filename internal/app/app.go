package app

import (
	"goSandbox/internal/tasks"
	"goSandbox/pkg/logger/zerolog"
)

func StartTasks() {
	logger := zerolog.LoggerBuilder("dev")

	tasks.First()
	logger.Info().Str("Task", "1").Msg("completed!")

	tasks.Second()
	logger.Info().Str("Task", "2").Msg("completed!")

	tasks.Third()
	logger.Info().Str("Task", "3").Msg("completed!")

	tasks.Fourth()
	logger.Info().Str("Task", "4").Msg("completed!")

	tasks.Fifth()
	logger.Info().Str("Task", "5").Msg("completed!")

	tasks.Sixth()
	logger.Info().Str("Task", "6").Msg("completed!")

	tasks.Seventh()
	logger.Info().Str("Task", "7").Msg("completed!")

	tasks.Eighth()
	logger.Info().Str("Task", "8").Msg("completed!")

	tasks.Ninth()
	logger.Info().Str("Task", "9").Msg("completed!")

	tasks.Tenth()
	logger.Info().Str("Task", "10").Msg("completed!")

	tasks.Eleventh()
	logger.Info().Str("Task", "11").Msg("completed!")

	tasks.Twelve()
	logger.Info().Str("Task", "12").Msg("completed!")

	tasks.Thirteenth()
	logger.Info().Str("Task", "13").Msg("completed!")

	tasks.Fourteenth()
	logger.Info().Str("Task", "14").Msg("completed!")

	tasks.Fifteenth()
	logger.Info().Str("Task", "15").Msg("completed!")

	tasks.Sixteenth()
	logger.Info().Str("Task", "16").Msg("completed!")

	tasks.Seventeenth()
	logger.Info().Str("Task", "17").Msg("completed!")

	tasks.Eighteenth()
	logger.Info().Str("Task", "18").Msg("completed!")

	tasks.Nineteenth()
	logger.Info().Str("Task", "19").Msg("completed!")

	tasks.Twentieth()
	logger.Info().Str("Task", "20").Msg("completed!")

	tasks.TwentyFirst()
	logger.Info().Str("Task", "21").Msg("completed!")

	tasks.TwentySecond()
	logger.Info().Str("Task", "22").Msg("completed!")

	tasks.TwentyThree()
	logger.Info().Str("Task", "23").Msg("completed!")

	tasks.TwentyFourth()
	logger.Info().Str("Task", "24").Msg("completed!")

	tasks.TwentyFifth()
	logger.Info().Str("Task", "25").Msg("completed!")

	tasks.TwentySixth()
	logger.Info().Str("Task", "26").Msg("completed!")
}
