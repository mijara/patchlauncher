package port

type PatcherService interface {
	Patch(romPath, patchPath, outputPath string) error
}
