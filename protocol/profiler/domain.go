// Code generated by cdpgen. DO NOT EDIT.

// Package profiler implements the Profiler domain.
package profiler

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Profiler domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Profiler domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Disable invokes the Profiler method.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Profiler method.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Enable", Err: err}
	}
	return
}

// GetBestEffortCoverage invokes the Profiler method. Collect coverage data
// for the current isolate. The coverage data may be incomplete due to garbage
// collection.
func (d *domainClient) GetBestEffortCoverage(ctx context.Context) (reply *GetBestEffortCoverageReply, err error) {
	reply = new(GetBestEffortCoverageReply)
	err = rpcc.Invoke(ctx, "Profiler.getBestEffortCoverage", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "GetBestEffortCoverage", Err: err}
	}
	return
}

// SetSamplingInterval invokes the Profiler method. Changes CPU profiler
// sampling interval. Must be called before CPU profiles recording started.
func (d *domainClient) SetSamplingInterval(ctx context.Context, args *SetSamplingIntervalArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Profiler.setSamplingInterval", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Profiler.setSamplingInterval", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "SetSamplingInterval", Err: err}
	}
	return
}

// Start invokes the Profiler method.
func (d *domainClient) Start(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.start", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Start", Err: err}
	}
	return
}

// StartPreciseCoverage invokes the Profiler method. Enable precise code
// coverage. Coverage data for JavaScript executed before enabling precise code
// coverage may be incomplete. Enabling prevents running optimized code and
// resets execution counters.
func (d *domainClient) StartPreciseCoverage(ctx context.Context, args *StartPreciseCoverageArgs) (reply *StartPreciseCoverageReply, err error) {
	reply = new(StartPreciseCoverageReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Profiler.startPreciseCoverage", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Profiler.startPreciseCoverage", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StartPreciseCoverage", Err: err}
	}
	return
}

// StartTypeProfile invokes the Profiler method. Enable type profile.
func (d *domainClient) StartTypeProfile(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.startTypeProfile", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StartTypeProfile", Err: err}
	}
	return
}

// Stop invokes the Profiler method.
func (d *domainClient) Stop(ctx context.Context) (reply *StopReply, err error) {
	reply = new(StopReply)
	err = rpcc.Invoke(ctx, "Profiler.stop", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Stop", Err: err}
	}
	return
}

// StopPreciseCoverage invokes the Profiler method. Disable precise code
// coverage. Disabling releases unnecessary execution count records and allows
// executing optimized code.
func (d *domainClient) StopPreciseCoverage(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.stopPreciseCoverage", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StopPreciseCoverage", Err: err}
	}
	return
}

// StopTypeProfile invokes the Profiler method. Disable type profile.
// Disabling releases type profile data collected so far.
func (d *domainClient) StopTypeProfile(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.stopTypeProfile", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StopTypeProfile", Err: err}
	}
	return
}

// TakePreciseCoverage invokes the Profiler method. Collect coverage data for
// the current isolate, and resets execution counters. Precise code coverage
// needs to have started.
func (d *domainClient) TakePreciseCoverage(ctx context.Context) (reply *TakePreciseCoverageReply, err error) {
	reply = new(TakePreciseCoverageReply)
	err = rpcc.Invoke(ctx, "Profiler.takePreciseCoverage", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "TakePreciseCoverage", Err: err}
	}
	return
}

// TakeTypeProfile invokes the Profiler method. Collect type profile.
func (d *domainClient) TakeTypeProfile(ctx context.Context) (reply *TakeTypeProfileReply, err error) {
	reply = new(TakeTypeProfileReply)
	err = rpcc.Invoke(ctx, "Profiler.takeTypeProfile", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "TakeTypeProfile", Err: err}
	}
	return
}

// EnableCounters invokes the Profiler method. Enable counters collection.
func (d *domainClient) EnableCounters(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.enableCounters", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "EnableCounters", Err: err}
	}
	return
}

// DisableCounters invokes the Profiler method. Disable counters collection.
func (d *domainClient) DisableCounters(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.disableCounters", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "DisableCounters", Err: err}
	}
	return
}

// GetCounters invokes the Profiler method. Retrieve counters.
func (d *domainClient) GetCounters(ctx context.Context) (reply *GetCountersReply, err error) {
	reply = new(GetCountersReply)
	err = rpcc.Invoke(ctx, "Profiler.getCounters", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "GetCounters", Err: err}
	}
	return
}

// EnableRuntimeCallStats invokes the Profiler method. Enable run time call
// stats collection.
func (d *domainClient) EnableRuntimeCallStats(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.enableRuntimeCallStats", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "EnableRuntimeCallStats", Err: err}
	}
	return
}

// DisableRuntimeCallStats invokes the Profiler method. Disable run time call
// stats collection.
func (d *domainClient) DisableRuntimeCallStats(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.disableRuntimeCallStats", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "DisableRuntimeCallStats", Err: err}
	}
	return
}

// GetRuntimeCallStats invokes the Profiler method. Retrieve run time call
// stats.
func (d *domainClient) GetRuntimeCallStats(ctx context.Context) (reply *GetRuntimeCallStatsReply, err error) {
	reply = new(GetRuntimeCallStatsReply)
	err = rpcc.Invoke(ctx, "Profiler.getRuntimeCallStats", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "GetRuntimeCallStats", Err: err}
	}
	return
}

func (d *domainClient) ConsoleProfileFinished(ctx context.Context) (ConsoleProfileFinishedClient, error) {
	s, err := rpcc.NewStream(ctx, "Profiler.consoleProfileFinished", d.conn)
	if err != nil {
		return nil, err
	}
	return &consoleProfileFinishedClient{Stream: s}, nil
}

type consoleProfileFinishedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *consoleProfileFinishedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *consoleProfileFinishedClient) Recv() (*ConsoleProfileFinishedReply, error) {
	event := new(ConsoleProfileFinishedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Profiler", Op: "ConsoleProfileFinished Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ConsoleProfileStarted(ctx context.Context) (ConsoleProfileStartedClient, error) {
	s, err := rpcc.NewStream(ctx, "Profiler.consoleProfileStarted", d.conn)
	if err != nil {
		return nil, err
	}
	return &consoleProfileStartedClient{Stream: s}, nil
}

type consoleProfileStartedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *consoleProfileStartedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *consoleProfileStartedClient) Recv() (*ConsoleProfileStartedReply, error) {
	event := new(ConsoleProfileStartedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Profiler", Op: "ConsoleProfileStarted Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) PreciseCoverageDeltaUpdate(ctx context.Context) (PreciseCoverageDeltaUpdateClient, error) {
	s, err := rpcc.NewStream(ctx, "Profiler.preciseCoverageDeltaUpdate", d.conn)
	if err != nil {
		return nil, err
	}
	return &preciseCoverageDeltaUpdateClient{Stream: s}, nil
}

type preciseCoverageDeltaUpdateClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *preciseCoverageDeltaUpdateClient) GetStream() rpcc.Stream { return c.Stream }

func (c *preciseCoverageDeltaUpdateClient) Recv() (*PreciseCoverageDeltaUpdateReply, error) {
	event := new(PreciseCoverageDeltaUpdateReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Profiler", Op: "PreciseCoverageDeltaUpdate Recv", Err: err}
	}
	return event, nil
}
