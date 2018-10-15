package main

import "math/rand"

var Categories = []string{
	"Audio/Video Clips",
	"Business/Economy",
	"Content Servers",
	"Email",
	"File Storage/Sharing",
	"Internet Telephony",
	"Mixed Content/Potentially Adult",
	"Newsgroups/Forums",
	"Non-Viewable/Infrastructure",
	"Reference",
	"Search Engines/Portals",
	"Shopping",
	"Social Networking",
	"Sports/Recreation",
	"Software Downloads",
	"Technology/Internet",
	"Travel",
	"Web Ads/Analytics",
}

func FakeCategories(r *rand.Rand) []string {
	categories := make([]string, 0, len(Categories))
	categories = append(categories, Categories...)
	nbCategories := int(r.ExpFloat64()*2.0 + 1.0)
	if nbCategories > len(categories) {
		nbCategories = len(categories)
	}
	r.Shuffle(len(categories), func(i, j int) { categories[i], categories[j] = categories[j], categories[i] })
	return categories[0:nbCategories]
}
