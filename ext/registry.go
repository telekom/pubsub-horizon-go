package ext

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Registry struct {
	extensions map[string]Extension
}

func NewRegistry() *Registry {
	return &Registry{
		extensions: make(map[string]Extension),
	}
}

func (r *Registry) RegisterDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".so") {
			continue
		}

		ext, err := open(filepath.Join(dir, entry.Name()))
		if err != nil {
			return err
		}

		if err := r.Register(ext); err != nil {
			return err
		}
	}

	return nil
}

func (r *Registry) Register(ext Extension) error {
	extensionName := ext.Info().Name
	if r.IsRegistered(ext) {
		return fmt.Errorf("theres already an extension with the name '%s'", extensionName)
	}

	r.extensions[extensionName] = ext
	return nil
}

func (r *Registry) IsRegistered(ext Extension) bool {
	extensionName := ext.Info().Name
	_, ok := r.extensions[extensionName]
	return ok
}

func (r *Registry) RunLifecyclePhase(phase LifecyclePhase) {
	for _, ext := range r.extensions {
		switch phase {

		case LifecycleConfigure:
			ext.Configure()

		case LifecycleEnable:
			ext.Enable()

		case LifecycleDisable:
			ext.Disable()

		default:
			panic(fmt.Errorf("unknown lifecycle phase with id '%d'", int(phase)))

		}
	}
}
