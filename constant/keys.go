package constant

type ContextKey string

const RepositoryKey ContextKey = "REPOSITORY"
const LoggerKey ContextKey = "LOGGER"
const LoaderKey ContextKey = "LOADER"

const (
	PlanetLoaderKey             string = "PLANET_LOADER"
	CharacterLoaderKey          string = "CHARACTER_LOADER"
	CharacterBySpeciesLoaderKey string = "CHARACTER_BY_SPECIES_LOADER"
	SpeciesLoaderKey            string = "SPECIES_LOADER"
)
