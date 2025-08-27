package seeders

import (
	"fmt"
	"log/slog"
)

type DatabaseSeeder interface {
	Seed() error
}

var seederRegistry []DatabaseSeeder

func RegisterSeeder(seeder DatabaseSeeder) {
	seederRegistry = append(seederRegistry, seeder)
}

func RegisterSeeders(seederList []DatabaseSeeder) {
	seederRegistry = append(seederRegistry, seederList...)
}

func GetSeeders() []DatabaseSeeder {
	return seederRegistry
}

func RunSeeders() error {
	slog.Info("=== Running Database Seeders ===")

	for _, seeder := range GetSeeders() {
		err := seeder.Seed()
		if err != nil {
			slog.Info("Error running seeder", slog.Any("error", err))
			return err
		}
	}

	fmt.Println("=== Database Seeding Completed Successfully ===")
	return nil
}
