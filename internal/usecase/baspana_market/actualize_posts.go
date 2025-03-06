package baspana

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	"github.com/google/uuid"
)

func (b *Baspana) actualizePosts(r *http.Request, p Post) (database.Postsbaspana, error){
	normolizedData, err := normolizePost(p)
	if err != nil {
		return database.Postsbaspana{}, fmt.Errorf("error on normalizing Post: %v", err)
	}
	_, err = b.DB.GetPost(r.Context(), database.GetPostParams{
		NumberObject: normolizedData.NumberObject, 
		LinkDetailPost: normolizedData.LinkDetailPost,
	})
	if err != nil {
		dbPostParams := database.CreatePostParams{
			ID:              normolizedData.ID,
			CreatedAt:       normolizedData.CreatedAt,
			UpdatedAt:       normolizedData.UpdatedAt,
			DatePublication: normolizedData.DatePublication,
			Title:           normolizedData.Title,
			Image:           normolizedData.Image,
			Address:         normolizedData.Address,
			CostForMetr:     normolizedData.CostForMetr,
			LinkDetailPost:  normolizedData.LinkDetailPost,
			NumberObject:    normolizedData.NumberObject,
			CountAccess:     normolizedData.CountAccess,
			Region:          normolizedData.Region,
		}
		dbPost, err := b.DB.CreatePost(r.Context(), dbPostParams)
		if err != nil {
			return database.Postsbaspana{}, fmt.Errorf("error creating post: %v", err)
		}

		return dbPost, nil
	}

	return database.Postsbaspana{}, fmt.Errorf("There is already a post with link: %v", normolizedData.LinkDetailPost)
}

func (b *Baspana) CreatePost(r *http.Request, n string) (database.Postsbaspana, error) {
	index, err := strconv.Atoi(n)
	if err != nil {
		return database.Postsbaspana{}, fmt.Errorf("error converting data: %v", err)
	}
	post := (*b.bufferPosts)[index]
	normolizedData, err := normolizePost(post)
	b.log.Info("normolizedData: %v", normolizedData)
	if err != nil {
		return database.Postsbaspana{}, err
	}
	dbPostParams := database.CreatePostParams{
		ID:              normolizedData.ID,
		CreatedAt:       normolizedData.CreatedAt,
		UpdatedAt:       normolizedData.UpdatedAt,
		DatePublication: normolizedData.DatePublication,
		Title:           normolizedData.Title,
		Image:           normolizedData.Image,
		Address:         normolizedData.Address,
		CostForMetr:     normolizedData.CostForMetr,
		LinkDetailPost:  normolizedData.LinkDetailPost,
		NumberObject:    normolizedData.NumberObject,
		CountAccess:     normolizedData.CountAccess,
		Region:          normolizedData.Region,
	}
	dbPost, err := b.DB.CreatePost(r.Context(), dbPostParams)
	if err != nil {
		return database.Postsbaspana{}, fmt.Errorf("error creating post: %v", err)
	}
	return dbPost, nil
}


func normolizePost(post Post) (database.Postsbaspana, error) {
	date, localTime, err := getNormalizedDate(post.Date)
	if err != nil {
		return database.Postsbaspana{}, fmt.Errorf("error Date on normalizing data: %v", err)
	}
	cost, err := strconv.Atoi(post.Cost)
	numberOfObject, err := strconv.Atoi(post.Id)
	count := 0
	if post.Count != "" {
		count, err = strconv.Atoi(post.Count)
	} 
	if err != nil {
		return database.Postsbaspana{}, fmt.Errorf(
			"error converting data: %v \nlink to post: %v, cost:%v, numberOfObject: %v, count: %v", 
			err, post.Link, cost, numberOfObject, count,
		)
	}
	dbPost := database.Postsbaspana{
		ID:              uuid.New(),
		CreatedAt:       localTime,
		UpdatedAt:       localTime,
		DatePublication: date,
		Title:           post.Title,
		Image:           post.Img,
		Address:         post.Address,
		CostForMetr:     int32(cost),
		LinkDetailPost:  post.Link,
		NumberObject:    int32(numberOfObject),
		CountAccess:     int32(count),
		Region:          getRegion(post.Address),
	}
	return dbPost, nil
}

func getNormalizedDate(dateStr string) (date time.Time, localTime time.Time,  err error){
	var normalizedDate time.Time
	utcPlus5 := time.FixedZone("UTC+5", 5*60*60)
	localTime = time.Now().UTC().In(utcPlus5)
	if (strings.ToLower(dateStr) == "сегодня" && dateStr == "") {
		normalizedDate = localTime
	} else {
		arrDate := strings.Split(dateStr, ".")
		day, err := strconv.Atoi(arrDate[0])
		month, err := strconv.Atoi(arrDate[1])
		year, err := strconv.Atoi(arrDate[2])
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("error converting date:", err)
		}
		normalizedDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}
	return normalizedDate, localTime, nil
}

func getRegion(address string) string{
	region := "none"
	mapRegions := map[string]string{
		"Западно-Казахстанская": "zko",
		"Астана": "astana",
		"Алматы": "almaty",
		"Шымкент": "shymkent",
		"область Абай": "abai",
		"Акмолинская": "akmola",
		"Актюбинская": "aktube",
		"Алматинская": "alm",
		"Атырауская": "atyrau",
		"Восточно-Казахстанская": "east",
		"Жамбылская": "zhanbyl",
		"Жетісу": "zhetysu",
		"Карагандинская": "karaganda",
		"Костанайская": "kostanay",
		"Кызылординская": "kysylorda",
		"Мангистауская": "mangistau",
		"Павлодарская": "pavlodar",
		"Северо-Казахстанская": "north",
		"Туркестанская": "turkestan",
	}
	for  key, value := range mapRegions {
		if (strings.Contains(strings.ToLower(address), strings.ToLower(key)) ||
			(strings.Contains(strings.ToLower(address), strings.ToLower("Алматы")) &&
				!strings.Contains(strings.ToLower(address), strings.ToLower("район Алматы"))) ||
				(strings.Contains(strings.ToLower(address), strings.ToLower("Астана")) &&
					!strings.Contains(strings.ToLower(address), strings.ToLower("район Астана")))) {
			region = value
		} 
	}
	return region
}