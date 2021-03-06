package calcium

import (
	"fmt"
	"strings"

	"github.com/projecteru2/core/cluster"
	"github.com/projecteru2/core/network"
	"github.com/projecteru2/core/network/sdn"
	"github.com/projecteru2/core/scheduler"
	"github.com/projecteru2/core/scheduler/complex"
	"github.com/projecteru2/core/source"
	"github.com/projecteru2/core/source/github"
	"github.com/projecteru2/core/source/gitlab"
	"github.com/projecteru2/core/store"
	"github.com/projecteru2/core/store/etcd"
	"github.com/projecteru2/core/types"
)

//Calcium implement the cluster
type Calcium struct {
	store     store.Store
	config    types.Config
	scheduler scheduler.Scheduler
	network   network.Network
	source    source.Source
}

// New returns a new cluster config
func New(config types.Config) (*Calcium, error) {
	// set store
	store, err := etcdstore.New(config)
	if err != nil {
		return nil, err
	}

	// set scheduler
	scheduler, err := complexscheduler.New(config)
	if err != nil {
		return nil, err
	}

	// set network
	titanium := sdn.New()

	// set scm
	var scm source.Source
	scmtype := strings.ToLower(config.Git.SCMType)
	switch scmtype {
	case cluster.GITLAB:
		scm = gitlab.New(config)
	case cluster.GITHUB:
		scm = github.New(config)
	default:
		return nil, fmt.Errorf("Unkonwn SCM type: %s", config.Git.SCMType)
	}

	return &Calcium{store: store, config: config, scheduler: scheduler, network: titanium, source: scm}, nil
}

//SetStore 用于在单元测试中更改etcd store为mock store
func (c *Calcium) SetStore(s store.Store) {
	c.store = s
}
