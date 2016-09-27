package context

type CancelFunc func()

type Ctx interface {
	OnCancel(data interface{}, callback func(ctx Ctx, data interface{}))
	Cancelled() bool

	cancel()
	addChild(ctx Ctx)
	removeChild(ctx Ctx)
}

type callback struct {
	f    func(ctx Ctx, data interface{})
	data interface{}
}

type ctx struct {
	parent    Ctx
	children  map[Ctx]bool
	callbacks []*callback
	cancelled bool
}

func (ctx *ctx) cancel() {
	if ctx.cancelled {
		// Can't cancel twice
		return
	}

	if ctx.parent != nil {
		ctx.parent.removeChild(ctx)
	}
	for child := range ctx.children {
		child.cancel()
	}
	for _, cb := range ctx.callbacks {
		cb.f(ctx, cb.data)
	}
	ctx.cancelled = true
}

func (ctx *ctx) OnCancel(data interface{}, f func(ctx Ctx, data interface{})) {
	// Call it immediately if we're already cancelled
	if ctx.cancelled {
		f(ctx, data)
		return
	}

	ctx.callbacks = append(ctx.callbacks, &callback{f, data})
}

func (ctx *ctx) Cancelled() bool {
	return ctx.cancelled
}

func (ctx *ctx) addChild(child Ctx) {
	ctx.children[child] = true
}

func (ctx *ctx) removeChild(child Ctx) {
	delete(ctx.children, child)
}

func New(parent Ctx) (Ctx, CancelFunc) {
	c := &ctx{parent: parent}
	parent.addChild(c)
	return c, c.cancel
}

type baseCtx struct{}

func (ctx *baseCtx) OnCancel(data interface{}, callback func(ctx Ctx, data interface{})) {}

func (ctx *baseCtx) Cancelled() bool { return false }

func (ctx *baseCtx) cancel() {}

func (ctx *baseCtx) addChild(child Ctx) {}

func (ctx *baseCtx) removeChild(child Ctx) {}

var Base = &baseCtx{}
