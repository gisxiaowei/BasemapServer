package conf

/*
 * conf.xml
 */

type CacheInfo struct {
	Typens           string `xml:"typens,attr"`
	TileCacheInfo    TileCacheInfo
	TileImageInfo    TileImageInfo
	CacheStorageInfo CacheStorageInfo
}

type TileCacheInfo struct {
	SpatialReference SpatialReference
	TileOrigin       TileOrigin
	TileCols         int64
	TileRows         int64
	DPI              int64
	PreciseDPI       int64
	LODInfos         []LODInfo `xml:"LODInfos>LODInfo"`
}

type TileImageInfo struct {
	CacheTileFormat    string
	CompressionQuality int64
	Antialiasing       bool
}

type CacheStorageInfo struct {
	StorageFormat string
	PacketSize    int64
}

type SpatialReference struct {
	WKT           string
	XOrigin       int64
	YOrigin       int64
	XYScale       float64
	ZOrigin       int64
	ZScale        int64
	MOrigin       int64
	MScale        int64
	XYTolerance   float64
	ZTolerance    float64
	MTolerance    float64
	HighPrecision bool
	LeftLongitude int64
	WKID          int64
	LatestWKID    int64
}

type TileOrigin struct {
	X int64
	Y int64
}

type LODInfo struct {
	LevelID    int64
	Scale      int64
	Resolution float64
}

/*
 * conf.cdi
 */

// EnvelopeN 矩形闭包
type EnvelopeN struct {
	XMin float64
	YMin float64
	XMax float64
	YMax float64
}
