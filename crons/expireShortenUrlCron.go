package crons

import (
	"goUrlShortener/repository/short_entity_repository"
)

func ExpireShortenUrlCronJob() {
	expiredShortenEntities := short_entity_repository.GetAllExpiredShortenEntities()
	for _, expiredShortenEntity := range expiredShortenEntities {
		short_entity_repository.Delete(&expiredShortenEntity)
	}
}
