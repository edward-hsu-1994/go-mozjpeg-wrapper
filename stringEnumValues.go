package go_mozjpeg_wrapper

type DCTMethod string
type TuneMode string
type QuantTable int

const (
	DCTInt   DCTMethod = "int"
	DCTFast  DCTMethod = "fast"
	DCTFloat DCTMethod = "float"

	TunePSNR    TuneMode = "psnr"
	TuneHVSPSNR TuneMode = "hvs-psnr"
	TuneSSIM    TuneMode = "ssim"
	TuneMSSSIM  TuneMode = "ms-ssim"
)

var validDCTMethods = map[DCTMethod]bool{
	DCTInt:   true,
	DCTFast:  true,
	DCTFloat: true,
}

var validTuneModes = map[TuneMode]bool{
	TunePSNR:    true,
	TuneHVSPSNR: true,
	TuneSSIM:    true,
	TuneMSSSIM:  true,
}

var validQuantTables = map[QuantTable]bool{
	0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true,
}
