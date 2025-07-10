package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed go_learning_roadmap.json
var roadmapData []byte

type Step struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Resource string `json:"resource"`
	Task     string `json:"task"`
}

type Progress struct {
	CurrentStep int `json:"current_step"`
}

func getProgressFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("❌ Cannot access user config dir:", err)
		os.Exit(1)
	}
	xgoDir := filepath.Join(configDir, "xgo")
	os.MkdirAll(xgoDir, 0755)
	return filepath.Join(xgoDir, "progress.json")
}

func loadRoadmap() ([]Step, error) {
	var steps []Step
	err := json.Unmarshal(roadmapData, &steps)
	return steps, err
}

func loadProgress() (*Progress, error) {
	path := getProgressFilePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Progress{CurrentStep: 0}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var progress Progress
	err = json.Unmarshal(data, &progress)
	return &progress, err
}

func saveProgress(p *Progress) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(getProgressFilePath(), data, 0644)
}

func showStep(step Step, index int) {
	fmt.Printf("\n📘 Step %d: %s\n", index+1, step.Title)
	fmt.Println("🕒 Duration:", step.Duration)
	fmt.Println("📚 Resource:", step.Resource)
	fmt.Println("✅ Task:", step.Task)
	fmt.Println()
}

func resetProgress() error {
	return saveProgress(&Progress{CurrentStep: 0})
}

func main() {
	now := flag.Bool("now", false, "Show the current learning step")
	done := flag.Bool("done", false, "Mark current step as done and move to next")
	reset := flag.Bool("reset", false, "Reset progress to step 1")

	flag.Parse()

	roadmap, err := loadRoadmap()
	if err != nil {
		fmt.Println("❌ Failed to load embedded roadmap:", err)
		return
	}

	progress, err := loadProgress()
	if err != nil {
		fmt.Println("❌ Failed to load progress:", err)
		return
	}

	switch {
	case *reset:
		err := resetProgress()
		if err != nil {
			fmt.Println("❌ Failed to reset progress:", err)
			return
		}
		fmt.Println("🔄 Progress has been reset to Step 1.")
	case *now:
		if progress.CurrentStep >= len(roadmap) {
			fmt.Println("🎉 All steps completed!")
			return
		}
		showStep(roadmap[progress.CurrentStep], progress.CurrentStep)
	case *done:
		if progress.CurrentStep >= len(roadmap) {
			fmt.Println("✅ All steps already done.")
			return
		}
		progress.CurrentStep++
		err := saveProgress(progress)
		if err != nil {
			fmt.Println("❌ Error saving progress:", err)
			return
		}
		if progress.CurrentStep < len(roadmap) {
			fmt.Println("✅ Progress saved. Next step:")
			showStep(roadmap[progress.CurrentStep], progress.CurrentStep)
		} else {
			fmt.Println("🎊 Congratulations! You’ve completed all steps.")
		}
	default:
		fmt.Println("Usage:\n  xgo -now     # Show current step\n  xgo -done    # Mark current step as done\n  xgo -reset   # Reset progress to step 1")
	}
}
