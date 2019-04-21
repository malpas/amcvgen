package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// NewCVFromText creates a new CV from text
func NewCVFromText(text string) (*CV, error) {
	cv := CV{}
	err := yaml.Unmarshal([]byte(text), &cv)

	if err != nil {
		return &cv, err
	}

	return &cv, err
}

// NewCVFromFile extracts resume data from a given file location
func NewCVFromFile(fileName string) *CV {
	fData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load %s. See -h\n", fileName)
		os.Exit(1)
	}
	cv, err := NewCVFromText(string(fData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse file: %s\n", err.Error())
		os.Exit(1)
	}
	return cv
}

// CV is the main CV data structure
// Uses the open JSON Resume schema (https://jsonresume.org/schema/)
type CV struct {
	Basics       CVBasics        `yaml:"basics"`
	Work         []CVWork        `yaml:"work"`
	Volunteer    []CVVolunteer   `yaml:"volunteer"`
	Education    []CVEducation   `yaml:"education"`
	Awards       []CVAward       `yaml:"awards"`
	Languages    []CVLanguage    `yaml:"language"`
	Interests    []CVInterest    `yaml:"interests"`
	References   []CVReference   `yaml:"references"`
	Skills       []CVSkill       `yaml:"skills"`
	Publications []CVPublication `yaml:"publications"`
}

// CVBasics contains general contact details and the resume image
type CVBasics struct {
	Name     string `yaml:"name"`
	Label    string `yaml:"label"`
	Picture  string `yaml:"picture"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	Website  string `yaml:"website"`
	Summary  string `yaml:"summary"`
	Location CVLocation
	Profiles []CVProfile
}

// CVLocation provides the applicant address
type CVLocation struct {
	Address     string
	PostalCode  string
	City        string
	CountryCode string
	Region      string
}

// CVProfile describes a social media account
type CVProfile struct {
	Network  string
	Username string
	URL      string
}

// CVWork provides information about a job
type CVWork struct {
	Company    string   `yaml:"company"`
	Position   string   `yaml:"position"`
	StartDate  string   `yaml:"startDate"`
	EndDate    string   `yaml:"endDate"`
	Summary    string   `yaml:"summary"`
	Highlights []string `yaml:"highlights"`
}

// CVVolunteer provides information about a volunteer work
type CVVolunteer struct {
	Organization string
	Position     string
	StartDate    string
	EndDate      string
	Summary      string
	Highlights   []string
}

// CVEducation provides information on formal education
type CVEducation struct {
	Institution string   `yaml:"institution"`
	Area        string   `yaml:"area"`
	StudyType   string   `yaml:"studyType"`
	StartDate   string   `yaml:"startDate"`
	EndDate     string   `yaml:"endDate"`
	GPA         string   `yaml:"gpa"`
	Courses     []string `yaml:"courses"`
}

// CVAward describes an award
type CVAward struct {
	Title   string
	Date    string
	Awarder string
	Summary string
}

// CVSkill describes a skill, with keywords and mastery
type CVSkill struct {
	Name     string
	Level    string
	Keywords []string
}

// CVInterest describes an interest with keywords
type CVInterest struct {
	Name     string
	Keywords []string
}

// CVLanguage describes a language proficiency
type CVLanguage struct {
	Language string
	Fluency  string
}

// CVReferences describes a reference
type CVReference struct {
	Name      string
	Reference string
}

// CVPublication describes a publication
type CVPublication struct {
	Name        string
	Publisher   string
	ReleaseDate string
	Website     string
	Summary     string
}
