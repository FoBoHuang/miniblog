package main

import "time"

// GolangCIConfig 代表 golangci-lint 的完整配置结构
type GolangCIConfig struct {
	Run             RunConfig             `yaml:"run" json:"run"`
	Output          OutputConfig          `yaml:"output" json:"output"`
	LintersSettings LintersSettingsConfig `yaml:"linters-settings" json:"linters-settings"`
	Linters         LintersConfig         `yaml:"linters" json:"linters"`
	Issues          IssuesConfig          `yaml:"issues" json:"issues"`
	Severity        SeverityConfig        `yaml:"severity" json:"severity"`
}

// RunConfig 运行配置
type RunConfig struct {
	Concurrency          int           `yaml:"concurrency" json:"concurrency"`
	Go                   string        `yaml:"go" json:"go"`
	Timeout              time.Duration `yaml:"timeout" json:"timeout"`
	IssuesExitCode       int           `yaml:"issues-exit-code" json:"issues-exit-code"`
	Tests                bool          `yaml:"tests" json:"tests"`
	BuildTags            []string      `yaml:"build-tags" json:"build-tags"`
	AllowParallelRunners bool          `yaml:"allow-parallel-runners" json:"allow-parallel-runners"`
}

// OutputConfig 输出配置
type OutputConfig struct {
	Formats          []string `yaml:"formats" json:"formats"`
	PrintIssuedLines bool     `yaml:"print-issued-lines" json:"print-issued-lines"`
	PrintLinterName  bool     `yaml:"print-linter-name" json:"print-linter-name"`
	PathPrefix       string   `yaml:"path-prefix" json:"path-prefix"`
	SortResults      bool     `yaml:"sort-results" json:"sort-results"`
}

// LintersSettingsConfig 代码检查器设置配置
type LintersSettingsConfig struct {
	Bidichk          BidichkConfig          `yaml:"bidichk" json:"bidichk"`
	Dogsled          DogsLedConfig          `yaml:"dogsled" json:"dogsled"`
	Dupl             DuplConfig             `yaml:"dupl" json:"dupl"`
	Errcheck         ErrcheckConfig         `yaml:"errcheck" json:"errcheck"`
	ErrorLint        ErrorLintConfig        `yaml:"errorlint" json:"errorlint"`
	Exhaustive       ExhaustiveConfig       `yaml:"exhaustive" json:"exhaustive"`
	ExhaustiveStruct ExhaustiveStructConfig `yaml:"exhaustivestruct" json:"exhaustivestruct"`
	Forbidigo        ForbidigConfig         `yaml:"forbidigo" json:"forbidigo"`
	FunLen           FunLenConfig           `yaml:"funlen" json:"funlen"`
	GoCognit         GoCognitConfig         `yaml:"gocognit" json:"gocognit"`
	GoConst          GoConstConfig          `yaml:"goconst" json:"goconst"`
	GoCyclo          GoCycloConfig          `yaml:"gocyclo" json:"gocyclo"`
	Cyclop           CyclopConfig           `yaml:"cyclop" json:"cyclop"`
	GoDot            GoDotConfig            `yaml:"godot" json:"godot"`
	GoDox            GoDoxConfig            `yaml:"godox" json:"godox"`
	GoFmt            GoFmtConfig            `yaml:"gofmt" json:"gofmt"`
	GoFumpt          GoFumptConfig          `yaml:"gofumpt" json:"gofumpt"`
	GoHeader         GoHeaderConfig         `yaml:"goheader" json:"goheader"`
	GoImports        GoImportsConfig        `yaml:"goimports" json:"goimports"`
	GoLint           GoLintConfig           `yaml:"golint" json:"golint"`
	GoMnd            GoMndConfig            `yaml:"gomnd" json:"gomnd"`
	GoModDirectives  GoModDirectivesConfig  `yaml:"gomoddirectives" json:"gomoddirectives"`
	GoModGuard       GoModGuardConfig       `yaml:"gomodguard" json:"gomodguard"`
	GoSec            GoSecConfig            `yaml:"gosec" json:"gosec"`
	GoSimple         GoSimpleConfig         `yaml:"gosimple" json:"gosimple"`
	GoVet            GoVetConfig            `yaml:"govet" json:"govet"`
	DepGuard         DepGuardConfig         `yaml:"depguard" json:"depguard"`
	IfShort          IfShortConfig          `yaml:"ifshort" json:"ifshort"`
	ImportAs         ImportAsConfig         `yaml:"importas" json:"importas"`
	IReturn          IReturnConfig          `yaml:"ireturn" json:"ireturn"`
	Lll              LllConfig              `yaml:"lll" json:"lll"`
	Maligned         MalignedConfig         `yaml:"maligned" json:"maligned"`
	Misspell         MisspellConfig         `yaml:"misspell" json:"misspell"`
	NakedRet         NakedRetConfig         `yaml:"nakedret" json:"nakedret"`
	NestIf           NestIfConfig           `yaml:"nestif" json:"nestif"`
	NilNil           NilNilConfig           `yaml:"nilnil" json:"nilnil"`
	NlReturn         NlReturnConfig         `yaml:"nlreturn" json:"nlreturn"`
	NoLintLint       NoLintLintConfig       `yaml:"nolintlint" json:"nolintlint"`
	Prealloc         PreallocConfig         `yaml:"prealloc" json:"prealloc"`
	PromLinter       PromLinterConfig       `yaml:"promlinter" json:"promlinter"`
	Predeclared      PredeclaredConfig      `yaml:"predeclared" json:"predeclared"`
	RowsErrCheck     RowsErrCheckConfig     `yaml:"rowserrcheck" json:"rowserrcheck"`
	Revive           ReviveConfig           `yaml:"revive" json:"revive"`
	StaticCheck      StaticCheckConfig      `yaml:"staticcheck" json:"staticcheck"`
	StyleCheck       StyleCheckConfig       `yaml:"stylecheck" json:"stylecheck"`
	Tagliatelle      TagliatelleConfig      `yaml:"tagliatelle" json:"tagliatelle"`
	TestPackage      TestPackageConfig      `yaml:"testpackage" json:"testpackage"`
	THelper          THelperConfig          `yaml:"thelper" json:"thelper"`
	TEnv             TEnvConfig             `yaml:"tenv" json:"tenv"`
	UnParam          UnParamConfig          `yaml:"unparam" json:"unparam"`
	Unused           UnusedConfig           `yaml:"unused" json:"unused"`
	Whitespace       WhitespaceConfig       `yaml:"whitespace" json:"whitespace"`
	WrapCheck        WrapCheckConfig        `yaml:"wrapcheck" json:"wrapcheck"`
	Wsl              WslConfig              `yaml:"wsl" json:"wsl"`
	MakeZero         MakeZeroConfig         `yaml:"makezero" json:"makezero"`
}

// BidichkConfig bidichk 检查器配置
type BidichkConfig struct {
	LeftToRightEmbedding     bool `yaml:"left-to-right-embedding" json:"left-to-right-embedding"`
	RightToLeftEmbedding     bool `yaml:"right-to-left-embedding" json:"right-to-left-embedding"`
	PopDirectionalFormatting bool `yaml:"pop-directional-formatting" json:"pop-directional-formatting"`
	LeftToRightOverride      bool `yaml:"left-to-right-override" json:"left-to-right-override"`
	RightToLeftOverride      bool `yaml:"right-to-left-override" json:"right-to-left-override"`
	LeftToRightIsolate       bool `yaml:"left-to-right-isolate" json:"left-to-right-isolate"`
	RightToLeftIsolate       bool `yaml:"right-to-left-isolate" json:"right-to-left-isolate"`
	FirstStrongIsolate       bool `yaml:"first-strong-isolate" json:"first-strong-isolate"`
	PopDirectionalIsolate    bool `yaml:"pop-directional-isolate" json:"pop-directional-isolate"`
}

// DogsLedConfig dogsled 检查器配置
type DogsLedConfig struct {
	MaxBlankIdentifiers int `yaml:"max-blank-identifiers" json:"max-blank-identifiers"`
}

// DuplConfig dupl 检查器配置
type DuplConfig struct {
	Threshold int `yaml:"threshold" json:"threshold"`
}

// ErrcheckConfig errcheck 检查器配置
type ErrcheckConfig struct {
	CheckTypeAssertions bool `yaml:"check-type-assertions" json:"check-type-assertions"`
	CheckBlank          bool `yaml:"check-blank" json:"check-blank"`
}

// ErrorLintConfig errorlint 检查器配置
type ErrorLintConfig struct {
	Errorf     bool `yaml:"errorf" json:"errorf"`
	Asserts    bool `yaml:"asserts" json:"asserts"`
	Comparison bool `yaml:"comparison" json:"comparison"`
}

// ExhaustiveConfig exhaustive 检查器配置
type ExhaustiveConfig struct {
	CheckGenerated             bool   `yaml:"check-generated" json:"check-generated"`
	DefaultSignifiesExhaustive bool   `yaml:"default-signifies-exhaustive" json:"default-signifies-exhaustive"`
	IgnoreEnumMembers          string `yaml:"ignore-enum-members" json:"ignore-enum-members"`
	PackageScopeOnly           bool   `yaml:"package-scope-only" json:"package-scope-only"`
}

// ExhaustiveStructConfig exhaustivestruct 检查器配置
type ExhaustiveStructConfig struct {
	StructPatterns []string `yaml:"struct-patterns" json:"struct-patterns"`
}

// ForbidigConfig forbidigo 检查器配置
type ForbidigConfig struct {
	Forbid               []string `yaml:"forbid" json:"forbid"`
	ExcludeGodocExamples bool     `yaml:"exclude_godoc_examples" json:"exclude_godoc_examples"`
}

// FunLenConfig funlen 检查器配置
type FunLenConfig struct {
	Lines      int `yaml:"lines" json:"lines"`
	Statements int `yaml:"statements" json:"statements"`
}

// GoCognitConfig gocognit 检查器配置
type GoCognitConfig struct {
	MinComplexity int `yaml:"min-complexity" json:"min-complexity"`
}

// GoConstConfig goconst 检查器配置
type GoConstConfig struct {
	MinLen         int  `yaml:"min-len" json:"min-len"`
	MinOccurrences int  `yaml:"min-occurrences" json:"min-occurrences"`
	IgnoreTests    bool `yaml:"ignore-tests" json:"ignore-tests"`
	MatchConstant  bool `yaml:"match-constant" json:"match-constant"`
	Numbers        bool `yaml:"numbers" json:"numbers"`
	Min            int  `yaml:"min" json:"min"`
	Max            int  `yaml:"max" json:"max"`
	IgnoreCalls    bool `yaml:"ignore-calls" json:"ignore-calls"`
}

// GoCycloConfig gocyclo 检查器配置
type GoCycloConfig struct {
	MinComplexity int `yaml:"min-complexity" json:"min-complexity"`
}

// CyclopConfig cyclop 检查器配置
type CyclopConfig struct {
	MaxComplexity  int     `yaml:"max-complexity" json:"max-complexity"`
	PackageAverage float64 `yaml:"package-average" json:"package-average"`
	SkipTests      bool    `yaml:"skip-tests" json:"skip-tests"`
}

// GoDotConfig godot 检查器配置
type GoDotConfig struct {
	Scope   string   `yaml:"scope" json:"scope"`
	Exclude []string `yaml:"exclude" json:"exclude"`
	Capital bool     `yaml:"capital" json:"capital"`
}

// GoDoxConfig godox 检查器配置
type GoDoxConfig struct {
	Keywords []string `yaml:"keywords" json:"keywords"`
}

// GoFmtConfig gofmt 检查器配置
type GoFmtConfig struct {
	Simplify bool `yaml:"simplify" json:"simplify"`
}

// GoFumptConfig gofumpt 检查器配置
type GoFumptConfig struct {
	ExtraRules bool `yaml:"extra-rules" json:"extra-rules"`
}

// GoHeaderConfig goheader 检查器配置
type GoHeaderConfig struct {
	Values       GoHeaderValues `yaml:"values" json:"values"`
	Template     *string        `yaml:"template" json:"template"`
	TemplatePath *string        `yaml:"template-path" json:"template-path"`
}

// GoHeaderValues goheader 值配置
type GoHeaderValues struct {
	Const  map[string]string `yaml:"const" json:"const"`
	Regexp map[string]string `yaml:"regexp" json:"regexp"`
}

// GoImportsConfig goimports 检查器配置
type GoImportsConfig struct {
	LocalPrefixes string `yaml:"local-prefixes" json:"local-prefixes"`
}

// GoLintConfig golint 检查器配置
type GoLintConfig struct {
	MinConfidence float64 `yaml:"min-confidence" json:"min-confidence"`
}

// GoMndConfig gomnd 检查器配置
type GoMndConfig struct {
	Settings GoMndSettings `yaml:"settings" json:"settings"`
}

// GoMndSettings gomnd 设置
type GoMndSettings struct {
	Mnd GoMndMnd `yaml:"mnd" json:"mnd"`
}

// GoMndMnd gomnd 魔法数字配置
type GoMndMnd struct {
	Checks string `yaml:"checks" json:"checks"`
}

// GoModDirectivesConfig gomoddirectives 检查器配置
type GoModDirectivesConfig struct {
	ReplaceLocal              bool     `yaml:"replace-local" json:"replace-local"`
	ReplaceAllowList          []string `yaml:"replace-allow-list" json:"replace-allow-list"`
	RetractAllowNoExplanation bool     `yaml:"retract-allow-no-explanation" json:"retract-allow-no-explanation"`
	ExcludeForbidden          bool     `yaml:"exclude-forbidden" json:"exclude-forbidden"`
}

// GoModGuardConfig gomodguard 检查器配置
type GoModGuardConfig struct {
	Allowed GoModGuardAllowed `yaml:"allowed" json:"allowed"`
	Blocked GoModGuardBlocked `yaml:"blocked" json:"blocked"`
}

// GoModGuardAllowed 允许的模块配置
type GoModGuardAllowed struct {
	Modules []string `yaml:"modules" json:"modules"`
	Domains []string `yaml:"domains" json:"domains"`
}

// GoModGuardBlocked 阻止的模块配置
type GoModGuardBlocked struct {
	Modules                []map[string]GoModGuardBlockedModule  `yaml:"modules" json:"modules"`
	Versions               []map[string]GoModGuardBlockedVersion `yaml:"versions" json:"versions"`
	LocalReplaceDirectives bool                                  `yaml:"local_replace_directives" json:"local_replace_directives"`
}

// GoModGuardBlockedModule 阻止的模块详情
type GoModGuardBlockedModule struct {
	Recommendations []string `yaml:"recommendations" json:"recommendations"`
	Reason          string   `yaml:"reason" json:"reason"`
}

// GoModGuardBlockedVersion 阻止的版本详情
type GoModGuardBlockedVersion struct {
	Version string `yaml:"version" json:"version"`
	Reason  string `yaml:"reason" json:"reason"`
}

// GoSecConfig gosec 检查器配置
type GoSecConfig struct {
	Includes         []string               `yaml:"includes" json:"includes"`
	Excludes         []string               `yaml:"excludes" json:"excludes"`
	ExcludeGenerated bool                   `yaml:"exclude-generated" json:"exclude-generated"`
	Severity         string                 `yaml:"severity" json:"severity"`
	Confidence       string                 `yaml:"confidence" json:"confidence"`
	Config           map[string]interface{} `yaml:"config" json:"config"`
}

// GoSimpleConfig gosimple 检查器配置
type GoSimpleConfig struct {
	Checks []string `yaml:"checks" json:"checks"`
}

// GoVetConfig govet 检查器配置
type GoVetConfig struct {
	DisableAll bool                              `yaml:"disable-all" json:"disable-all"`
	EnableAll  bool                              `yaml:"enable-all" json:"enable-all"`
	Disable    []string                          `yaml:"disable" json:"disable"`
	Settings   map[string]map[string]interface{} `yaml:"settings" json:"settings"`
}

// DepGuardConfig depguard 检查器配置
type DepGuardConfig struct {
	Rules map[string]DepGuardRule `yaml:"rules" json:"rules"`
}

// DepGuardRule depguard 规则
type DepGuardRule struct {
	ListMode string                `yaml:"list-mode" json:"list-mode"`
	Files    []string              `yaml:"files" json:"files"`
	Deny     []DepGuardDenyPackage `yaml:"deny" json:"deny"`
}

// DepGuardDenyPackage depguard 拒绝的包
type DepGuardDenyPackage struct {
	Pkg  string `yaml:"pkg" json:"pkg"`
	Desc string `yaml:"desc" json:"desc"`
}

// IfShortConfig ifshort 检查器配置
type IfShortConfig struct {
	MaxDeclLines int `yaml:"max-decl-lines" json:"max-decl-lines"`
	MaxDeclChars int `yaml:"max-decl-chars" json:"max-decl-chars"`
}

// ImportAsConfig importas 检查器配置
type ImportAsConfig struct {
	NoUnaliased bool                `yaml:"no-unaliased" json:"no-unaliased"`
	Alias       []ImportAsAliasRule `yaml:"alias" json:"alias"`
	JWT         string              `yaml:"jwt" json:"jwt"`
	MetaV1      string              `yaml:"metav1" json:"metav1"`
}

// ImportAsAliasRule importas 别名规则
type ImportAsAliasRule struct {
	Pkg   string `yaml:"pkg" json:"pkg"`
	Alias string `yaml:"alias" json:"alias"`
}

// IReturnConfig ireturn 检查器配置
type IReturnConfig struct {
	Allow  []string `yaml:"allow" json:"allow"`
	Reject []string `yaml:"reject" json:"reject"`
}

// LllConfig lll 检查器配置
type LllConfig struct {
	LineLength int `yaml:"line-length" json:"line-length"`
	TabWidth   int `yaml:"tab-width" json:"tab-width"`
}

// MalignedConfig maligned 检查器配置
type MalignedConfig struct {
	SuggestNew bool `yaml:"suggest-new" json:"suggest-new"`
}

// MisspellConfig misspell 检查器配置
type MisspellConfig struct {
	Locale      string   `yaml:"locale" json:"locale"`
	IgnoreWords []string `yaml:"ignore-words" json:"ignore-words"`
}

// NakedRetConfig nakedret 检查器配置
type NakedRetConfig struct {
	MaxFuncLines int `yaml:"max-func-lines" json:"max-func-lines"`
}

// NestIfConfig nestif 检查器配置
type NestIfConfig struct {
	MinComplexity int `yaml:"min-complexity" json:"min-complexity"`
}

// NilNilConfig nilnil 检查器配置
type NilNilConfig struct {
	CheckedTypes []string `yaml:"checked-types" json:"checked-types"`
}

// NlReturnConfig nlreturn 检查器配置
type NlReturnConfig struct {
	BlockSize int `yaml:"block-size" json:"block-size"`
}

// NoLintLintConfig nolintlint 检查器配置
type NoLintLintConfig struct {
	AllowUnused        bool     `yaml:"allow-unused" json:"allow-unused"`
	AllowLeadingSpace  bool     `yaml:"allow-leading-space" json:"allow-leading-space"`
	AllowNoExplanation []string `yaml:"allow-no-explanation" json:"allow-no-explanation"`
	RequireExplanation bool     `yaml:"require-explanation" json:"require-explanation"`
	RequireSpecific    bool     `yaml:"require-specific" json:"require-specific"`
}

// PreallocConfig prealloc 检查器配置
type PreallocConfig struct {
	Simple     bool `yaml:"simple" json:"simple"`
	RangeLoops bool `yaml:"range-loops" json:"range-loops"`
	ForLoops   bool `yaml:"for-loops" json:"for-loops"`
}

// PromLinterConfig promlinter 检查器配置
type PromLinterConfig struct {
	Strict          bool     `yaml:"strict" json:"strict"`
	DisabledLinters []string `yaml:"disabled-linters" json:"disabled-linters"`
}

// PredeclaredConfig predeclared 检查器配置
type PredeclaredConfig struct {
	Ignore string `yaml:"ignore" json:"ignore"`
	Q      bool   `yaml:"q" json:"q"`
}

// RowsErrCheckConfig rowserrcheck 检查器配置
type RowsErrCheckConfig struct {
	Packages []string `yaml:"packages" json:"packages"`
}

// ReviveConfig revive 检查器配置
type ReviveConfig struct {
	IgnoreGeneratedHeader bool         `yaml:"ignore-generated-header" json:"ignore-generated-header"`
	Severity              string       `yaml:"severity" json:"severity"`
	Rules                 []ReviveRule `yaml:"rules" json:"rules"`
}

// ReviveRule revive 规则
type ReviveRule struct {
	Name     string `yaml:"name" json:"name"`
	Severity string `yaml:"severity" json:"severity"`
}

// StaticCheckConfig staticcheck 检查器配置
type StaticCheckConfig struct {
	Checks []string `yaml:"checks" json:"checks"`
}

// StyleCheckConfig stylecheck 检查器配置
type StyleCheckConfig struct {
	Checks                  []string `yaml:"checks" json:"checks"`
	DotImportWhitelist      []string `yaml:"dot-import-whitelist" json:"dot-import-whitelist"`
	Initialisms             []string `yaml:"initialisms" json:"initialisms"`
	HTTPStatusCodeWhitelist []string `yaml:"http-status-code-whitelist" json:"http-status-code-whitelist"`
}

// TagliatelleConfig tagliatelle 检查器配置
type TagliatelleConfig struct {
	Case TagliatelleCase `yaml:"case" json:"case"`
}

// TagliatelleCase tagliatelle 大小写配置
type TagliatelleCase struct {
	UseFieldName bool              `yaml:"use-field-name" json:"use-field-name"`
	Rules        map[string]string `yaml:"rules" json:"rules"`
}

// TestPackageConfig testpackage 检查器配置
type TestPackageConfig struct {
	SkipRegexp string `yaml:"skip-regexp" json:"skip-regexp"`
}

// THelperConfig thelper 检查器配置
type THelperConfig struct {
	Test      THelperTestConfig `yaml:"test" json:"test"`
	Benchmark THelperTestConfig `yaml:"benchmark" json:"benchmark"`
	TB        THelperTestConfig `yaml:"tb" json:"tb"`
}

// THelperTestConfig thelper 测试配置
type THelperTestConfig struct {
	First bool `yaml:"first" json:"first"`
	Name  bool `yaml:"name" json:"name"`
	Begin bool `yaml:"begin" json:"begin"`
}

// TEnvConfig tenv 检查器配置
type TEnvConfig struct {
	All bool `yaml:"all" json:"all"`
}

// UnParamConfig unparam 检查器配置
type UnParamConfig struct {
	CheckExported bool `yaml:"check-exported" json:"check-exported"`
}

// UnusedConfig unused 检查器配置
type UnusedConfig struct {
	CheckExported bool `yaml:"check-exported" json:"check-exported"`
}

// WhitespaceConfig whitespace 检查器配置
type WhitespaceConfig struct {
	MultiIf   bool `yaml:"multi-if" json:"multi-if"`
	MultiFunc bool `yaml:"multi-func" json:"multi-func"`
}

// WrapCheckConfig wrapcheck 检查器配置
type WrapCheckConfig struct {
	IgnoreSigs         []string `yaml:"ignoreSigs" json:"ignoreSigs"`
	IgnorePackageGlobs []string `yaml:"ignorePackageGlobs" json:"ignorePackageGlobs"`
}

// WslConfig wsl 检查器配置
type WslConfig struct {
	StrictAppend                 bool `yaml:"strict-append" json:"strict-append"`
	AllowAssignAndCall           bool `yaml:"allow-assign-and-call" json:"allow-assign-and-call"`
	AllowAssignAndAnything       bool `yaml:"allow-assign-and-anything" json:"allow-assign-and-anything"`
	AllowMultilineAssign         bool `yaml:"allow-multiline-assign" json:"allow-multiline-assign"`
	AllowCuddleDeclarations      bool `yaml:"allow-cuddle-declarations" json:"allow-cuddle-declarations"`
	AllowTrailingComment         bool `yaml:"allow-trailing-comment" json:"allow-trailing-comment"`
	ForceCaseTrailingWhitespace  int  `yaml:"force-case-trailing-whitespace" json:"force-case-trailing-whitespace"`
	ForceErrCuddling             bool `yaml:"force-err-cuddling" json:"force-err-cuddling"`
	AllowSeparatedLeadingComment bool `yaml:"allow-separated-leading-comment" json:"allow-separated-leading-comment"`
}

// MakeZeroConfig makezero 检查器配置
type MakeZeroConfig struct {
	Always bool `yaml:"always" json:"always"`
}

// LintersConfig linters 配置
type LintersConfig struct {
	EnableAll  bool     `yaml:"enable-all" json:"enable-all"`
	DisableAll bool     `yaml:"disable-all" json:"disable-all"`
	Disable    []string `yaml:"disable" json:"disable"`
	Fast       bool     `yaml:"fast" json:"fast"`
}

// IssuesConfig issues 配置
type IssuesConfig struct {
	Exclude               []string            `yaml:"exclude" json:"exclude"`
	ExcludeFiles          []string            `yaml:"exclude-files" json:"exclude-files"`
	ExcludeDirs           []string            `yaml:"exclude-dirs" json:"exclude-dirs"`
	ExcludeRules          []IssuesExcludeRule `yaml:"exclude-rules" json:"exclude-rules"`
	ExcludeUseDefault     bool                `yaml:"exclude-use-default" json:"exclude-use-default"`
	ExcludeDirsUseDefault bool                `yaml:"exclude-dirs-use-defaul" json:"exclude-dirs-use-defaul"`
	ExcludeCaseSensitive  bool                `yaml:"exclude-case-sensitive" json:"exclude-case-sensitive"`
	Include               []string            `yaml:"include" json:"include"`
	MaxIssuesPerLinter    int                 `yaml:"max-issues-per-linter" json:"max-issues-per-linter"`
	MaxSameIssues         int                 `yaml:"max-same-issues" json:"max-same-issues"`
	UniqByLine            bool                `yaml:"uniq-by-line" json:"uniq-by-line"`
	New                   bool                `yaml:"new" json:"new"`
	Fix                   bool                `yaml:"fix" json:"fix"`
}

// IssuesExcludeRule issues 排除规则
type IssuesExcludeRule struct {
	Linters []string `yaml:"linters" json:"linters"`
	Text    string   `yaml:"text,omitempty" json:"text,omitempty"`
	Source  string   `yaml:"source,omitempty" json:"source,omitempty"`
}

// SeverityConfig severity 配置
type SeverityConfig struct {
	DefaultSeverity string         `yaml:"default-severity" json:"default-severity"`
	CaseSensitive   bool           `yaml:"case-sensitive" json:"case-sensitive"`
	Rules           []SeverityRule `yaml:"rules" json:"rules"`
}

// SeverityRule severity 规则
type SeverityRule struct {
	Linters  []string `yaml:"linters" json:"linters"`
	Severity string   `yaml:"severity" json:"severity"`
}
