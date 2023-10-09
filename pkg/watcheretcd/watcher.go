package watcheretcd

import (
	"context"
	"runtime"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Watcher struct {
	// lock for callback
	lock        sync.RWMutex
	endpoints   []string
	client      *clientv3.Client
	running     bool
	callback    func(string)
	keyName     string
	password    string
	lastSentRev int64
}

func finalizer(w *Watcher) {
	w.running = false
}

func NewWatcher(endpoints []string, keyName string, password ...string) (*Watcher, error) {
	w := &Watcher{}
	w.endpoints = endpoints
	w.running = true
	w.callback = nil
	w.keyName = keyName
	if len(password) > 0 {
		w.password = password[0]
	}

	// Create the client.
	err := w.createClient()
	if err != nil {
		return nil, err
	}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(w, finalizer)

	go func() {
		_ = w.startWatch()
	}()

	return w, nil
}

// Close closes the Watcher.
func (w *Watcher) Close() {
	finalizer(w)
}

func (w *Watcher) createClient() error {
	cfg := clientv3.Config{
		Endpoints: w.endpoints,
		// set timeout per request to fail fast when the target endpoints is unavailable
		DialKeepAliveTimeout: time.Second * 10,
		DialTimeout:          time.Second * 30,
		Password:             w.password,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	w.client = c
	return nil
}

// SetUpdateCallback sets the callback function that the watcher will call
// when the policy in DB has been changed by other instances.
// A classic callback is Enforcer.LoadPolicy().
func (w *Watcher) SetUpdateCallback(callback func(string)) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.callback = callback
	return nil
}

func (w *Watcher) Update() error {
	w.lock.Lock()
	defer w.lock.Unlock()
	resp, err := w.client.Put(context.TODO(), w.keyName, "")
	if err == nil {
		w.lastSentRev = resp.Header.GetRevision()
	}
	return err
}

// startWatch is a goroutine that watches the policy change.
func (w *Watcher) startWatch() error {
	watcher := w.client.Watch(context.Background(), w.keyName)
	for res := range watcher {
		t := res.Events[0]
		if t.IsCreate() || t.IsModify() {
			w.lock.RLock()
			//ignore self update
			if rev := t.Kv.ModRevision; rev > w.lastSentRev && w.callback != nil {
				// w.callback(strconv.FormatInt(rev, 10))
				w.callback(string(t.Kv.Value))
			}
			w.lock.RUnlock()
		}
	}
	return nil
}
