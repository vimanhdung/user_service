package random

import (
	"awesomeProject/pb"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomGender() pb.User_Gender {
	switch rand.Intn(2) {
	case 1:
		return pb.User_FEMALE
	default:
		return pb.User_MALE
	}
}

func randomBirthday() time.Time {
	min := 1940
	max := 2000
	birthYear := min + rand.Intn(max - min)
	return time.Date(birthYear, 1, 0,0,0,0,0,time.Local)
}

func randomRole() int {
	return randomInt(1, 3);
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomRank() *pb.Rank {
	return &pb.Rank{
		Rank:      float32(randomFloat(1, 6)),
		Language:  randomLanguageLevel(),
		Framework: randomFrameworkLevel(),
	}
}

func randomLanguageLevel() []*pb.Language {
	numberLang := randomInt(1, 9)
	langs := make([]*pb.Language, 10)
	for i := 0; i < numberLang; i++ {
		langs = append(langs, &pb.Language{
			Language: randomLanguage(),
			Level:    float32(randomFloat(1, 5)),
			Years:    int32(randomInt(1, 20)),
		})
	}
	return langs
}

func randomLanguage() pb.Language_Lang {
	switch rand.Intn(8) {
	case 0:
		return pb.Language_PHP
	case 1:
		return pb.Language_JAVA
	case 2:
		return pb.Language_GO
	case 3:
		return pb.Language_PYTHON
	case 4:
		return pb.Language_RUBY
	case 5:
		return pb.Language_JAVASCRIPT
	case 6:
		return pb.Language_C_PLUS_PLUS
	case 7:
		return pb.Language_SCALA
	default:
		return pb.Language_SWIFT
	}
}

func randomFrameworkLevel() []*pb.FrameWork {
	numberFW := randomInt(1, 11)
	frameWorks := make([]*pb.FrameWork, 11)
	for i := 0; i < numberFW; i++ {
		frameWorks = append(frameWorks, &pb.FrameWork{
			Framwork: randomFramework(),
			Level:    float32(randomFloat(1, 5)),
			Years:    int32(randomFloat(1, 20)),
		})
	}
	return frameWorks
}

func randomFramework() pb.FrameWork_FW {
	switch rand.Intn(11) {
	case 0:
		return pb.FrameWork_RUBY_ON_RAILS
	case 1:
		return pb.FrameWork_SYMPHONY
	case 2:
		return pb.FrameWork_ANGULAR_JS
	case 3:
		return pb.FrameWork_REACT_JS
	case 4:
		return pb.FrameWork_CAKE
	case 5:
		return pb.FrameWork_NODE
	case 6:
		return pb.FrameWork_YII
	case 7:
		return pb.FrameWork_LARAVEL
	case 8:
		return pb.FrameWork_MAGENTO
	case 9:
		return pb.FrameWork_SPRING
	default:
		return pb.FrameWork_DRUPAL
	}
}