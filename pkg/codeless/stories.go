package codeless

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed stories.json
var storiesJson string

type Story struct {
	Id         int
	CaseNumber string
	Title      string
	Text       string
}

var stories []Story
var storiesById map[int]Story
var storyIds []int

func init() {
	err := json.Unmarshal([]byte(storiesJson), &stories)

	if err != nil {
		panic(err)
	}

	storiesById = make(map[int]Story)

	for i := 0; i < len(stories); i++ {
		if stories[i].CaseNumber == "" {
			continue
		}
		split := strings.Split(stories[i].CaseNumber, " ")
		idString := split[1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println(err)
		}
		storyIds = append(storyIds, id)
		stories[i].Id = id
		storiesById[id] = stories[i]
	}

	sort.Ints(storyIds)
}

func GetStoryById(id int) *Story {
	story := storiesById[id]
	return &story
}

func DecorateStory(story *Story) string {
	return fmt.Sprintf("# %s: %s\n\n%s", story.CaseNumber, story.Title, story.Text)
}

func GetAllStories() []Story {
	var sortedStories []Story

	for i := 0; i < len(storyIds); i++ {
		story := storiesById[i]
		if story.CaseNumber == "" {
			continue
		}
		sortedStories = append(sortedStories, story)
	}

	return sortedStories
}

func GetNextStoryId(id int) int {
	currentIndex := findIntIndex(id, storyIds)
	nextIndex := currentIndex + 1
	nextStoryId := storyIds[nextIndex]
	return storiesById[nextStoryId].Id
}

func GetRandomStory() *Story {
	rand.Seed(time.Now().UnixNano())
	max := len(storyIds)
	id := rand.Intn(max)
	story := GetStoryById(id)

	return story
}

func findIntIndex(value int, list []int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}
