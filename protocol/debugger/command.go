// Code generated by cdpgen. DO NOT EDIT.

package debugger

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// ContinueToLocationArgs represents the arguments for ContinueToLocation in the Debugger domain.
type ContinueToLocationArgs struct {
	Location Location `json:"location"` // Location to continue to.
	// TargetCallFrames
	//
	// Values: "any", "current".
	TargetCallFrames *string `json:"targetCallFrames,omitempty"`
}

// NewContinueToLocationArgs initializes ContinueToLocationArgs with the required arguments.
func NewContinueToLocationArgs(location Location) *ContinueToLocationArgs {
	args := new(ContinueToLocationArgs)
	args.Location = location
	return args
}

// SetTargetCallFrames sets the TargetCallFrames optional argument.
//
// Values: "any", "current".
func (a *ContinueToLocationArgs) SetTargetCallFrames(targetCallFrames string) *ContinueToLocationArgs {
	a.TargetCallFrames = &targetCallFrames
	return a
}

// EnableArgs represents the arguments for Enable in the Debugger domain.
type EnableArgs struct {
	// MaxScriptsCacheSize The maximum size in bytes of collected scripts
	// (not referenced by other heap objects) the debugger can hold. Puts
	// no limit if parameter is omitted.
	//
	// Note: This property is experimental.
	MaxScriptsCacheSize *float64 `json:"maxScriptsCacheSize,omitempty"`
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetMaxScriptsCacheSize sets the MaxScriptsCacheSize optional argument.
// The maximum size in bytes of collected scripts (not referenced by
// other heap objects) the debugger can hold. Puts no limit if
// parameter is omitted.
//
// Note: This property is experimental.
func (a *EnableArgs) SetMaxScriptsCacheSize(maxScriptsCacheSize float64) *EnableArgs {
	a.MaxScriptsCacheSize = &maxScriptsCacheSize
	return a
}

// EnableReply represents the return values for Enable in the Debugger domain.
type EnableReply struct {
	// DebuggerID Unique identifier of the debugger.
	//
	// Note: This property is experimental.
	DebuggerID runtime.UniqueDebuggerID `json:"debuggerId"`
}

// EvaluateOnCallFrameArgs represents the arguments for EvaluateOnCallFrame in the Debugger domain.
type EvaluateOnCallFrameArgs struct {
	CallFrameID           CallFrameID `json:"callFrameId"`                     // Call frame identifier to evaluate on.
	Expression            string      `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           *string     `json:"objectGroup,omitempty"`           // String object group name to put result into (allows rapid releasing resulting object handles using `releaseObjectGroup`).
	IncludeCommandLineAPI *bool       `json:"includeCommandLineAPI,omitempty"` // Specifies whether command line API should be available to the evaluated expression, defaults to false.
	Silent                *bool       `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	ReturnByValue         *bool       `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	// GeneratePreview Whether preview should be generated for the result.
	//
	// Note: This property is experimental.
	GeneratePreview   *bool `json:"generatePreview,omitempty"`
	ThrowOnSideEffect *bool `json:"throwOnSideEffect,omitempty"` // Whether to throw an exception if side effect cannot be ruled out during evaluation.
	// Timeout Terminate execution after timing out (number of
	// milliseconds).
	//
	// Note: This property is experimental.
	Timeout *runtime.TimeDelta `json:"timeout,omitempty"`
}

// NewEvaluateOnCallFrameArgs initializes EvaluateOnCallFrameArgs with the required arguments.
func NewEvaluateOnCallFrameArgs(callFrameID CallFrameID, expression string) *EvaluateOnCallFrameArgs {
	args := new(EvaluateOnCallFrameArgs)
	args.CallFrameID = callFrameID
	args.Expression = expression
	return args
}

// SetObjectGroup sets the ObjectGroup optional argument. String
// object group name to put result into (allows rapid releasing
// resulting object handles using `releaseObjectGroup`).
func (a *EvaluateOnCallFrameArgs) SetObjectGroup(objectGroup string) *EvaluateOnCallFrameArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// SetIncludeCommandLineAPI sets the IncludeCommandLineAPI optional argument.
// Specifies whether command line API should be available to the
// evaluated expression, defaults to false.
func (a *EvaluateOnCallFrameArgs) SetIncludeCommandLineAPI(includeCommandLineAPI bool) *EvaluateOnCallFrameArgs {
	a.IncludeCommandLineAPI = &includeCommandLineAPI
	return a
}

// SetSilent sets the Silent optional argument. In silent mode
// exceptions thrown during evaluation are not reported and do not
// pause execution. Overrides `setPauseOnException` state.
func (a *EvaluateOnCallFrameArgs) SetSilent(silent bool) *EvaluateOnCallFrameArgs {
	a.Silent = &silent
	return a
}

// SetReturnByValue sets the ReturnByValue optional argument. Whether
// the result is expected to be a JSON object that should be sent by
// value.
func (a *EvaluateOnCallFrameArgs) SetReturnByValue(returnByValue bool) *EvaluateOnCallFrameArgs {
	a.ReturnByValue = &returnByValue
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument.
// Whether preview should be generated for the result.
//
// Note: This property is experimental.
func (a *EvaluateOnCallFrameArgs) SetGeneratePreview(generatePreview bool) *EvaluateOnCallFrameArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// SetThrowOnSideEffect sets the ThrowOnSideEffect optional argument.
// Whether to throw an exception if side effect cannot be ruled out
// during evaluation.
func (a *EvaluateOnCallFrameArgs) SetThrowOnSideEffect(throwOnSideEffect bool) *EvaluateOnCallFrameArgs {
	a.ThrowOnSideEffect = &throwOnSideEffect
	return a
}

// SetTimeout sets the Timeout optional argument. Terminate execution
// after timing out (number of milliseconds).
//
// Note: This property is experimental.
func (a *EvaluateOnCallFrameArgs) SetTimeout(timeout runtime.TimeDelta) *EvaluateOnCallFrameArgs {
	a.Timeout = &timeout
	return a
}

// EvaluateOnCallFrameReply represents the return values for EvaluateOnCallFrame in the Debugger domain.
type EvaluateOnCallFrameReply struct {
	Result           runtime.RemoteObject      `json:"result"`                     // Object wrapper for the evaluation result.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// GetPossibleBreakpointsArgs represents the arguments for GetPossibleBreakpoints in the Debugger domain.
type GetPossibleBreakpointsArgs struct {
	Start              Location  `json:"start"`                        // Start of range to search possible breakpoint locations in.
	End                *Location `json:"end,omitempty"`                // End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	RestrictToFunction *bool     `json:"restrictToFunction,omitempty"` // Only consider locations which are in the same (non-nested) function as start.
}

// NewGetPossibleBreakpointsArgs initializes GetPossibleBreakpointsArgs with the required arguments.
func NewGetPossibleBreakpointsArgs(start Location) *GetPossibleBreakpointsArgs {
	args := new(GetPossibleBreakpointsArgs)
	args.Start = start
	return args
}

// SetEnd sets the End optional argument. End of range to search
// possible breakpoint locations in (excluding). When not specified,
// end of scripts is used as end of range.
func (a *GetPossibleBreakpointsArgs) SetEnd(end Location) *GetPossibleBreakpointsArgs {
	a.End = &end
	return a
}

// SetRestrictToFunction sets the RestrictToFunction optional argument.
// Only consider locations which are in the same (non-nested) function
// as start.
func (a *GetPossibleBreakpointsArgs) SetRestrictToFunction(restrictToFunction bool) *GetPossibleBreakpointsArgs {
	a.RestrictToFunction = &restrictToFunction
	return a
}

// GetPossibleBreakpointsReply represents the return values for GetPossibleBreakpoints in the Debugger domain.
type GetPossibleBreakpointsReply struct {
	Locations []BreakLocation `json:"locations"` // List of the possible breakpoint locations.
}

// GetScriptSourceArgs represents the arguments for GetScriptSource in the Debugger domain.
type GetScriptSourceArgs struct {
	ScriptID runtime.ScriptID `json:"scriptId"` // Id of the script to get source for.
}

// NewGetScriptSourceArgs initializes GetScriptSourceArgs with the required arguments.
func NewGetScriptSourceArgs(scriptID runtime.ScriptID) *GetScriptSourceArgs {
	args := new(GetScriptSourceArgs)
	args.ScriptID = scriptID
	return args
}

// GetScriptSourceReply represents the return values for GetScriptSource in the Debugger domain.
type GetScriptSourceReply struct {
	ScriptSource string `json:"scriptSource"`       // Script source (empty in case of Wasm bytecode).
	Bytecode     []byte `json:"bytecode,omitempty"` // Wasm bytecode. (Encoded as a base64 string when passed over JSON)
}

// GetWasmBytecodeArgs represents the arguments for GetWasmBytecode in the Debugger domain.
type GetWasmBytecodeArgs struct {
	ScriptID runtime.ScriptID `json:"scriptId"` // Id of the Wasm script to get source for.
}

// NewGetWasmBytecodeArgs initializes GetWasmBytecodeArgs with the required arguments.
func NewGetWasmBytecodeArgs(scriptID runtime.ScriptID) *GetWasmBytecodeArgs {
	args := new(GetWasmBytecodeArgs)
	args.ScriptID = scriptID
	return args
}

// GetWasmBytecodeReply represents the return values for GetWasmBytecode in the Debugger domain.
type GetWasmBytecodeReply struct {
	Bytecode []byte `json:"bytecode"` // Script source. (Encoded as a base64 string when passed over JSON)
}

// GetStackTraceArgs represents the arguments for GetStackTrace in the Debugger domain.
type GetStackTraceArgs struct {
	StackTraceID runtime.StackTraceID `json:"stackTraceId"` // No description.
}

// NewGetStackTraceArgs initializes GetStackTraceArgs with the required arguments.
func NewGetStackTraceArgs(stackTraceID runtime.StackTraceID) *GetStackTraceArgs {
	args := new(GetStackTraceArgs)
	args.StackTraceID = stackTraceID
	return args
}

// GetStackTraceReply represents the return values for GetStackTrace in the Debugger domain.
type GetStackTraceReply struct {
	StackTrace runtime.StackTrace `json:"stackTrace"` // No description.
}

// PauseOnAsyncCallArgs represents the arguments for PauseOnAsyncCall in the Debugger domain.
type PauseOnAsyncCallArgs struct {
	ParentStackTraceID runtime.StackTraceID `json:"parentStackTraceId"` // Debugger will pause when async call with given stack trace is started.
}

// NewPauseOnAsyncCallArgs initializes PauseOnAsyncCallArgs with the required arguments.
func NewPauseOnAsyncCallArgs(parentStackTraceID runtime.StackTraceID) *PauseOnAsyncCallArgs {
	args := new(PauseOnAsyncCallArgs)
	args.ParentStackTraceID = parentStackTraceID
	return args
}

// RemoveBreakpointArgs represents the arguments for RemoveBreakpoint in the Debugger domain.
type RemoveBreakpointArgs struct {
	BreakpointID BreakpointID `json:"breakpointId"` // No description.
}

// NewRemoveBreakpointArgs initializes RemoveBreakpointArgs with the required arguments.
func NewRemoveBreakpointArgs(breakpointID BreakpointID) *RemoveBreakpointArgs {
	args := new(RemoveBreakpointArgs)
	args.BreakpointID = breakpointID
	return args
}

// RestartFrameArgs represents the arguments for RestartFrame in the Debugger domain.
type RestartFrameArgs struct {
	CallFrameID CallFrameID `json:"callFrameId"` // Call frame identifier to evaluate on.
}

// NewRestartFrameArgs initializes RestartFrameArgs with the required arguments.
func NewRestartFrameArgs(callFrameID CallFrameID) *RestartFrameArgs {
	args := new(RestartFrameArgs)
	args.CallFrameID = callFrameID
	return args
}

// RestartFrameReply represents the return values for RestartFrame in the Debugger domain.
type RestartFrameReply struct {
	CallFrames      []CallFrame         `json:"callFrames"`                // New stack trace.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	// AsyncStackTraceID Async stack trace, if any.
	//
	// Note: This property is experimental.
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
}

// ResumeArgs represents the arguments for Resume in the Debugger domain.
type ResumeArgs struct {
	TerminateOnResume *bool `json:"terminateOnResume,omitempty"` // Set to true to terminate execution upon resuming execution. In contrast to Runtime.terminateExecution, this will allows to execute further JavaScript (i.e. via evaluation) until execution of the paused code is actually resumed, at which point termination is triggered. If execution is currently not paused, this parameter has no effect.
}

// NewResumeArgs initializes ResumeArgs with the required arguments.
func NewResumeArgs() *ResumeArgs {
	args := new(ResumeArgs)

	return args
}

// SetTerminateOnResume sets the TerminateOnResume optional argument.
// Set to true to terminate execution upon resuming execution. In
// contrast to Runtime.terminateExecution, this will allows to execute
// further JavaScript (i.e. via evaluation) until execution of the
// paused code is actually resumed, at which point termination is
// triggered. If execution is currently not paused, this parameter has
// no effect.
func (a *ResumeArgs) SetTerminateOnResume(terminateOnResume bool) *ResumeArgs {
	a.TerminateOnResume = &terminateOnResume
	return a
}

// SearchInContentArgs represents the arguments for SearchInContent in the Debugger domain.
type SearchInContentArgs struct {
	ScriptID      runtime.ScriptID `json:"scriptId"`                // Id of the script to search in.
	Query         string           `json:"query"`                   // String to search for.
	CaseSensitive *bool            `json:"caseSensitive,omitempty"` // If true, search is case sensitive.
	IsRegex       *bool            `json:"isRegex,omitempty"`       // If true, treats string parameter as regex.
}

// NewSearchInContentArgs initializes SearchInContentArgs with the required arguments.
func NewSearchInContentArgs(scriptID runtime.ScriptID, query string) *SearchInContentArgs {
	args := new(SearchInContentArgs)
	args.ScriptID = scriptID
	args.Query = query
	return args
}

// SetCaseSensitive sets the CaseSensitive optional argument. If true,
// search is case sensitive.
func (a *SearchInContentArgs) SetCaseSensitive(caseSensitive bool) *SearchInContentArgs {
	a.CaseSensitive = &caseSensitive
	return a
}

// SetIsRegex sets the IsRegex optional argument. If true, treats
// string parameter as regex.
func (a *SearchInContentArgs) SetIsRegex(isRegex bool) *SearchInContentArgs {
	a.IsRegex = &isRegex
	return a
}

// SearchInContentReply represents the return values for SearchInContent in the Debugger domain.
type SearchInContentReply struct {
	Result []SearchMatch `json:"result"` // List of search matches.
}

// SetAsyncCallStackDepthArgs represents the arguments for SetAsyncCallStackDepth in the Debugger domain.
type SetAsyncCallStackDepthArgs struct {
	MaxDepth int `json:"maxDepth"` // Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
}

// NewSetAsyncCallStackDepthArgs initializes SetAsyncCallStackDepthArgs with the required arguments.
func NewSetAsyncCallStackDepthArgs(maxDepth int) *SetAsyncCallStackDepthArgs {
	args := new(SetAsyncCallStackDepthArgs)
	args.MaxDepth = maxDepth
	return args
}

// SetBlackboxPatternsArgs represents the arguments for SetBlackboxPatterns in the Debugger domain.
type SetBlackboxPatternsArgs struct {
	Patterns []string `json:"patterns"` // Array of regexps that will be used to check script url for blackbox state.
}

// NewSetBlackboxPatternsArgs initializes SetBlackboxPatternsArgs with the required arguments.
func NewSetBlackboxPatternsArgs(patterns []string) *SetBlackboxPatternsArgs {
	args := new(SetBlackboxPatternsArgs)
	args.Patterns = patterns
	return args
}

// SetBlackboxedRangesArgs represents the arguments for SetBlackboxedRanges in the Debugger domain.
type SetBlackboxedRangesArgs struct {
	ScriptID  runtime.ScriptID `json:"scriptId"`  // Id of the script.
	Positions []ScriptPosition `json:"positions"` // No description.
}

// NewSetBlackboxedRangesArgs initializes SetBlackboxedRangesArgs with the required arguments.
func NewSetBlackboxedRangesArgs(scriptID runtime.ScriptID, positions []ScriptPosition) *SetBlackboxedRangesArgs {
	args := new(SetBlackboxedRangesArgs)
	args.ScriptID = scriptID
	args.Positions = positions
	return args
}

// SetBreakpointArgs represents the arguments for SetBreakpoint in the Debugger domain.
type SetBreakpointArgs struct {
	Location  Location `json:"location"`            // Location to set breakpoint in.
	Condition *string  `json:"condition,omitempty"` // Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

// NewSetBreakpointArgs initializes SetBreakpointArgs with the required arguments.
func NewSetBreakpointArgs(location Location) *SetBreakpointArgs {
	args := new(SetBreakpointArgs)
	args.Location = location
	return args
}

// SetCondition sets the Condition optional argument. Expression to
// use as a breakpoint condition. When specified, debugger will only
// stop on the breakpoint if this expression evaluates to true.
func (a *SetBreakpointArgs) SetCondition(condition string) *SetBreakpointArgs {
	a.Condition = &condition
	return a
}

// SetBreakpointReply represents the return values for SetBreakpoint in the Debugger domain.
type SetBreakpointReply struct {
	BreakpointID   BreakpointID `json:"breakpointId"`   // Id of the created breakpoint for further reference.
	ActualLocation Location     `json:"actualLocation"` // Location this breakpoint resolved into.
}

// SetInstrumentationBreakpointArgs represents the arguments for SetInstrumentationBreakpoint in the Debugger domain.
type SetInstrumentationBreakpointArgs struct {
	// Instrumentation Instrumentation name.
	//
	// Values: "beforeScriptExecution", "beforeScriptWithSourceMapExecution".
	Instrumentation string `json:"instrumentation"`
}

// NewSetInstrumentationBreakpointArgs initializes SetInstrumentationBreakpointArgs with the required arguments.
func NewSetInstrumentationBreakpointArgs(instrumentation string) *SetInstrumentationBreakpointArgs {
	args := new(SetInstrumentationBreakpointArgs)
	args.Instrumentation = instrumentation
	return args
}

// SetInstrumentationBreakpointReply represents the return values for SetInstrumentationBreakpoint in the Debugger domain.
type SetInstrumentationBreakpointReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Id of the created breakpoint for further reference.
}

// SetBreakpointByURLArgs represents the arguments for SetBreakpointByURL in the Debugger domain.
type SetBreakpointByURLArgs struct {
	LineNumber   int     `json:"lineNumber"`             // Line number to set breakpoint at.
	URL          *string `json:"url,omitempty"`          // URL of the resources to set breakpoint on.
	URLRegex     *string `json:"urlRegex,omitempty"`     // Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or `urlRegex` must be specified.
	ScriptHash   *string `json:"scriptHash,omitempty"`   // Script hash of the resources to set breakpoint on.
	ColumnNumber *int    `json:"columnNumber,omitempty"` // Offset in the line to set breakpoint at.
	Condition    *string `json:"condition,omitempty"`    // Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

// NewSetBreakpointByURLArgs initializes SetBreakpointByURLArgs with the required arguments.
func NewSetBreakpointByURLArgs(lineNumber int) *SetBreakpointByURLArgs {
	args := new(SetBreakpointByURLArgs)
	args.LineNumber = lineNumber
	return args
}

// SetURL sets the URL optional argument. URL of the resources to set
// breakpoint on.
func (a *SetBreakpointByURLArgs) SetURL(url string) *SetBreakpointByURLArgs {
	a.URL = &url
	return a
}

// SetURLRegex sets the URLRegex optional argument. Regex pattern for
// the URLs of the resources to set breakpoints on. Either `url` or
// `urlRegex` must be specified.
func (a *SetBreakpointByURLArgs) SetURLRegex(urlRegex string) *SetBreakpointByURLArgs {
	a.URLRegex = &urlRegex
	return a
}

// SetScriptHash sets the ScriptHash optional argument. Script hash of
// the resources to set breakpoint on.
func (a *SetBreakpointByURLArgs) SetScriptHash(scriptHash string) *SetBreakpointByURLArgs {
	a.ScriptHash = &scriptHash
	return a
}

// SetColumnNumber sets the ColumnNumber optional argument. Offset in
// the line to set breakpoint at.
func (a *SetBreakpointByURLArgs) SetColumnNumber(columnNumber int) *SetBreakpointByURLArgs {
	a.ColumnNumber = &columnNumber
	return a
}

// SetCondition sets the Condition optional argument. Expression to
// use as a breakpoint condition. When specified, debugger will only
// stop on the breakpoint if this expression evaluates to true.
func (a *SetBreakpointByURLArgs) SetCondition(condition string) *SetBreakpointByURLArgs {
	a.Condition = &condition
	return a
}

// SetBreakpointByURLReply represents the return values for SetBreakpointByURL in the Debugger domain.
type SetBreakpointByURLReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Id of the created breakpoint for further reference.
	Locations    []Location   `json:"locations"`    // List of the locations this breakpoint resolved into upon addition.
}

// SetBreakpointOnFunctionCallArgs represents the arguments for SetBreakpointOnFunctionCall in the Debugger domain.
type SetBreakpointOnFunctionCallArgs struct {
	ObjectID  runtime.RemoteObjectID `json:"objectId"`            // Function object id.
	Condition *string                `json:"condition,omitempty"` // Expression to use as a breakpoint condition. When specified, debugger will stop on the breakpoint if this expression evaluates to true.
}

// NewSetBreakpointOnFunctionCallArgs initializes SetBreakpointOnFunctionCallArgs with the required arguments.
func NewSetBreakpointOnFunctionCallArgs(objectID runtime.RemoteObjectID) *SetBreakpointOnFunctionCallArgs {
	args := new(SetBreakpointOnFunctionCallArgs)
	args.ObjectID = objectID
	return args
}

// SetCondition sets the Condition optional argument. Expression to
// use as a breakpoint condition. When specified, debugger will stop on
// the breakpoint if this expression evaluates to true.
func (a *SetBreakpointOnFunctionCallArgs) SetCondition(condition string) *SetBreakpointOnFunctionCallArgs {
	a.Condition = &condition
	return a
}

// SetBreakpointOnFunctionCallReply represents the return values for SetBreakpointOnFunctionCall in the Debugger domain.
type SetBreakpointOnFunctionCallReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Id of the created breakpoint for further reference.
}

// SetBreakpointsActiveArgs represents the arguments for SetBreakpointsActive in the Debugger domain.
type SetBreakpointsActiveArgs struct {
	Active bool `json:"active"` // New value for breakpoints active state.
}

// NewSetBreakpointsActiveArgs initializes SetBreakpointsActiveArgs with the required arguments.
func NewSetBreakpointsActiveArgs(active bool) *SetBreakpointsActiveArgs {
	args := new(SetBreakpointsActiveArgs)
	args.Active = active
	return args
}

// SetPauseOnExceptionsArgs represents the arguments for SetPauseOnExceptions in the Debugger domain.
type SetPauseOnExceptionsArgs struct {
	// State Pause on exceptions mode.
	//
	// Values: "none", "uncaught", "all".
	State string `json:"state"`
}

// NewSetPauseOnExceptionsArgs initializes SetPauseOnExceptionsArgs with the required arguments.
func NewSetPauseOnExceptionsArgs(state string) *SetPauseOnExceptionsArgs {
	args := new(SetPauseOnExceptionsArgs)
	args.State = state
	return args
}

// SetReturnValueArgs represents the arguments for SetReturnValue in the Debugger domain.
type SetReturnValueArgs struct {
	NewValue runtime.CallArgument `json:"newValue"` // New return value.
}

// NewSetReturnValueArgs initializes SetReturnValueArgs with the required arguments.
func NewSetReturnValueArgs(newValue runtime.CallArgument) *SetReturnValueArgs {
	args := new(SetReturnValueArgs)
	args.NewValue = newValue
	return args
}

// SetScriptSourceArgs represents the arguments for SetScriptSource in the Debugger domain.
type SetScriptSourceArgs struct {
	ScriptID     runtime.ScriptID `json:"scriptId"`         // Id of the script to edit.
	ScriptSource string           `json:"scriptSource"`     // New content of the script.
	DryRun       *bool            `json:"dryRun,omitempty"` // If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
}

// NewSetScriptSourceArgs initializes SetScriptSourceArgs with the required arguments.
func NewSetScriptSourceArgs(scriptID runtime.ScriptID, scriptSource string) *SetScriptSourceArgs {
	args := new(SetScriptSourceArgs)
	args.ScriptID = scriptID
	args.ScriptSource = scriptSource
	return args
}

// SetDryRun sets the DryRun optional argument. If true the change
// will not actually be applied. Dry run may be used to get result
// description without actually modifying the code.
func (a *SetScriptSourceArgs) SetDryRun(dryRun bool) *SetScriptSourceArgs {
	a.DryRun = &dryRun
	return a
}

// SetScriptSourceReply represents the return values for SetScriptSource in the Debugger domain.
type SetScriptSourceReply struct {
	CallFrames      []CallFrame         `json:"callFrames,omitempty"`      // New stack trace in case editing has happened while VM was stopped.
	StackChanged    *bool               `json:"stackChanged,omitempty"`    // Whether current call stack was modified after applying the changes.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	// AsyncStackTraceID Async stack trace, if any.
	//
	// Note: This property is experimental.
	AsyncStackTraceID *runtime.StackTraceID     `json:"asyncStackTraceId,omitempty"`
	ExceptionDetails  *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details if any.
}

// SetSkipAllPausesArgs represents the arguments for SetSkipAllPauses in the Debugger domain.
type SetSkipAllPausesArgs struct {
	Skip bool `json:"skip"` // New value for skip pauses state.
}

// NewSetSkipAllPausesArgs initializes SetSkipAllPausesArgs with the required arguments.
func NewSetSkipAllPausesArgs(skip bool) *SetSkipAllPausesArgs {
	args := new(SetSkipAllPausesArgs)
	args.Skip = skip
	return args
}

// SetVariableValueArgs represents the arguments for SetVariableValue in the Debugger domain.
type SetVariableValueArgs struct {
	ScopeNumber  int                  `json:"scopeNumber"`  // 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	VariableName string               `json:"variableName"` // Variable name.
	NewValue     runtime.CallArgument `json:"newValue"`     // New variable value.
	CallFrameID  CallFrameID          `json:"callFrameId"`  // Id of callframe that holds variable.
}

// NewSetVariableValueArgs initializes SetVariableValueArgs with the required arguments.
func NewSetVariableValueArgs(scopeNumber int, variableName string, newValue runtime.CallArgument, callFrameID CallFrameID) *SetVariableValueArgs {
	args := new(SetVariableValueArgs)
	args.ScopeNumber = scopeNumber
	args.VariableName = variableName
	args.NewValue = newValue
	args.CallFrameID = callFrameID
	return args
}

// StepIntoArgs represents the arguments for StepInto in the Debugger domain.
type StepIntoArgs struct {
	// BreakOnAsyncCall Debugger will pause on the execution of the first
	// async task which was scheduled before next pause.
	//
	// Note: This property is experimental.
	BreakOnAsyncCall *bool `json:"breakOnAsyncCall,omitempty"`
	// SkipList The skipList specifies location ranges that should be
	// skipped on step into.
	//
	// Note: This property is experimental.
	SkipList []LocationRange `json:"skipList,omitempty"`
}

// NewStepIntoArgs initializes StepIntoArgs with the required arguments.
func NewStepIntoArgs() *StepIntoArgs {
	args := new(StepIntoArgs)

	return args
}

// SetBreakOnAsyncCall sets the BreakOnAsyncCall optional argument.
// Debugger will pause on the execution of the first async task which
// was scheduled before next pause.
//
// Note: This property is experimental.
func (a *StepIntoArgs) SetBreakOnAsyncCall(breakOnAsyncCall bool) *StepIntoArgs {
	a.BreakOnAsyncCall = &breakOnAsyncCall
	return a
}

// SetSkipList sets the SkipList optional argument. The skipList
// specifies location ranges that should be skipped on step into.
//
// Note: This property is experimental.
func (a *StepIntoArgs) SetSkipList(skipList []LocationRange) *StepIntoArgs {
	a.SkipList = skipList
	return a
}

// StepOverArgs represents the arguments for StepOver in the Debugger domain.
type StepOverArgs struct {
	// SkipList The skipList specifies location ranges that should be
	// skipped on step over.
	//
	// Note: This property is experimental.
	SkipList []LocationRange `json:"skipList,omitempty"`
}

// NewStepOverArgs initializes StepOverArgs with the required arguments.
func NewStepOverArgs() *StepOverArgs {
	args := new(StepOverArgs)

	return args
}

// SetSkipList sets the SkipList optional argument. The skipList
// specifies location ranges that should be skipped on step over.
//
// Note: This property is experimental.
func (a *StepOverArgs) SetSkipList(skipList []LocationRange) *StepOverArgs {
	a.SkipList = skipList
	return a
}
