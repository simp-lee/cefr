# Requirements: simp-lee/cefr — Go CEFR 文本难度评估库

## Overview

`simp-lee/cefr` 是一个纯 Go 的英文文本 CEFR（Common European Framework of Reference）难度级别评估库。采用多特征融合算法（词汇分析 50% + 句法复杂度 30% + 可读性公式 20%），无需外部 NLP 服务，可离线对任意英文文本输出 A1-C2 级别评估和置信度评分。

## Intent Classification
- **Type**: New Feature（从零创建独立库）
- **Complexity**: High
- **Scope Confidence**: Medium — 算法参数需实际语料验证调优

## Goals

- **G1**: 对任意英文文本（句子/段落/文章/书籍章节）输出 CEFR 级别（A1-C2）和连续分数（1.0-6.0）
- **G2**: 评估结果包含各维度子分数（词汇/句法/可读性）和置信度，支持调用方理解评估依据
- **G3**: 内嵌词频数据（Oxford 5000 / NGSL / AWL + 不规则变体表），零外部依赖，可离线运行
- **G4**: 万字级文本评估 < 50ms（不含 I/O），十万字级文本通过采样策略保证性能
- **G5**: API 简洁，一个函数调用即可完成评估

## Non-Goals (Explicit Exclusions)

- **NG1**: 不做完整 NLP（不含词性标注、依存句法分析、命名实体识别等）
- **NG2**: 不做非英语文本评估（仅英语）
- **NG3**: 不做 CEFR **教学/训练**相关功能（仅评估）
- **NG4**: 不做语法错误检测/纠正
- **NG5**: 不做基于 LLM 的评估（纯规则 + 统计方法；LLM 辅助在调用方层面按需集成）
- **NG6**: 不做词典功能（义项级别 CEFR 标注由 `simp-lee/isdict-api` 提供）

---

## Functional Requirements

### FR-1xx: 核心评估

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-101 | **一站式评估 API**：接受纯文本字符串，返回 `Result` 结构体（CEFR 级别、连续分数、各维度子分数、置信度、词汇统计） | API 易用性 | `cefr.Assess("The cat sat on the mat.")` 返回 `Result{Level: "A1", Score: 1.2, ...}` |
| FR-102 | **多特征融合算法**：最终分数 = 词汇分数 × 0.50 + 句法分数 × 0.30 + 可读性分数 × 0.20。每个子分数输出 1.0-6.0 连续值 | Vajjala & Meurers (2012) 多特征方法 | 三个子分数在 [1.0, 6.0] 范围内；最终分数为加权平均 |
| FR-103 | **分数到级别映射**：1.0-1.5 → A1，1.5-2.5 → A2，2.5-3.5 → B1，3.5-4.5 → B2，4.5-5.5 → C1，5.5-6.0 → C2 | 标准映射 | `Score: 3.7` 映射为 `Level: "B2"` |
| FR-104 | **置信度评分**：基于文本长度和样本内部一致性输出 0.0-1.0 的置信度。极短文本（< 50 词）/采样段间差异大时置信度低 | 避免对极短文本过度相信评估结果 | 10 词文本置信度 < 0.5；1000 词文本置信度 > 0.8 |
| FR-105 | **可配置权重**：支持通过 Option 模式调整三维度权重 | 不同应用场景权重需求不同 | `cefr.Assess(text, cefr.WithWeights(0.6, 0.2, 0.2))` 生效 |

### FR-2xx: 词汇分析（权重 0.50）

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-201 | **内嵌 CEFR 词表数据**：将 Oxford 5000（含 CEFR A1-C1 标注）通过 `embed.FS` 编译到二进制中 | 零外部依赖的核心保障 | 库初始化不需要任何外部文件 |
| FR-202 | **分层词汇查找**：① Oxford 5000（精确 CEFR 标注）→ ② AWL 570 词族（映射 B2-C1）→ ③ NGSL 2800 词（按频率排名映射）→ ④ 未找到标记为 unknown（假设 C1） | 研究报告推荐的层级策略 | "elaborate" 在 Oxford 5000 中找到 B2；"albeit" 在 AWL 中找到 C1；"cat" 在 NGSL 中找到 A1 |
| FR-203 | **词形还原（Lemmatization）**：将变体还原为词典基本形（running→run, better→good, children→child）。采用不规则变体表（~1000 词）+ 规则后缀剥离 | 确保变体词汇能在词表中查找到 | "running" → "run" (A1)；"studied" → "study" (A2)；"children" → "child" (A1) |
| FR-204 | **内嵌不规则变体表**：维护常见不规则动词（went→go）、不规则名词（mice→mouse）、不规则形容词（better→good）的映射表 | 规则后缀剥离无法处理不规则变化 | "went" → "go"；"better" → "good" |
| FR-205 | **规则后缀剥离**：处理 `-ing`（双写还原：running→run、去e还原：making→make）、`-ed`（ied→y：studied→study）、`-s/-es`、`-er/-est`、`-ly` | 覆盖规则变化 | "walking" → "walk"；"happier" → "happy" |
| FR-206 | **停用词过滤**：内嵌 ~200 个英语停用词（the, a, is, in, on...），分析时跳过不计入级别 | 停用词不携带难度信息 | "The cat is on the mat" 仅 "cat" 和 "mat" 参与级别计算 |
| FR-207 | **专有名词检测（启发式）**：非句首且首字母大写 → 跳过；全大写（NASA）→ 跳过；连续大写开头词（New York）→ 跳过 | 专有名词不应影响难度评估 | "John went to London" 中 "John" 和 "London" 被跳过 |
| FR-208 | **数字和标点过滤**：纯数字和标点不参与分析 | 无难度信息 | "The year 2024 was..." 中 "2024" 被跳过 |
| FR-209 | **词汇分数计算（加权百分位法）**：将所有内容词的 CEFR 级别排序，取 P80（第 80 百分位）作为基础分数；未知词比例 > 5% 时向上修正（每超 1% 加 0.1） | 文本难度由较难词汇决定而非平均值 | 含 20% B2 词汇的文本得分高于含 5% B2 词汇的文本 |
| FR-210 | **词汇统计输出**：返回各级别词汇数量和百分比分布、未知词比例、总内容词数 | 调用方需要理解评估依据 | `result.Vocab.Distribution` 返回 `{A1: 45%, A2: 25%, B1: 15%, B2: 10%, C1: 3%, Unknown: 2%}` |

### FR-3xx: 句法复杂度分析（权重 0.30）

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-301 | **分句**：按 `.` `!` `?` 分句，处理缩写（Mr. Dr. U.S. e.g. i.e.）、省略号（...）、引号内句号不作为分句点 | 正确的句数是所有句法指标的基础 | "Mr. Smith went to Washington. He liked it!" → 2 句 |
| FR-302 | **平均句长（ASL）**：总词数 / 句数。映射函数：≤6→1.0, ≤9→2.0, ≤13→3.0, ≤18→4.0, ≤23→5.0, >23→6.0（线性插值） | 最强单一句法预测因子（Vajjala & Meurers 2012）| 平均句长 12 词的文本句法 ASL 子分 ≈ 2.8 |
| FR-303 | **从属指数**：统计从属连词/关系代词（because, although, which, who, if, when, while, unless, although, though, since, that, where, whose, whom 等 ~30 个标记词）出现次数 / 句数 | 区分 B1/B2 的关键指标 | 含大量从句的学术文本从属指数 > 1.0 |
| FR-304 | **被动语态检测（启发式）**：检测 `be动词 + 过去分词` bigrams。be 动词形式：is/am/are/was/were/been/being/be；过去分词：-ed 结尾或在不规则过去分词表中。统计被动频率 = 被动次数 / 句数 | 区分 A/B 到 B/C 级别 | "The book was written by..." 检测为被动语态 |
| FR-305 | **连接词多样性**：统计文本中出现的**不同类型**连接词数量（去重）。映射：0-2→A1, 3-5→A2, 6-10→B1, 11-15→B2, 16-20→C1, 20+→C2 | 区分中高级文本 | 使用了 because, although, however, furthermore, nevertheless 的文本多样性分 ≈ B2 |
| FR-306 | **句法综合分数**：`syntaxScore = 0.40×f(ASL) + 0.30×f(subordinationIndex) + 0.15×f(passiveRate) + 0.15×f(connectorDiversity)` | 多指标融合更稳定 | 各子指标权重可配置 |

### FR-4xx: 可读性公式（权重 0.20）

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-401 | **Flesch-Kincaid Grade Level (FKGL)**：$0.39 \times \frac{words}{sentences} + 11.8 \times \frac{syllables}{words} - 15.59$。输出美国年级（1-16+） | 最广泛使用的可读性公式 | "The cat sat on the mat." FKGL ≈ 1.0 |
| FR-402 | **Flesch Reading Ease (FRE)**：$206.835 - 1.015 \times \frac{words}{sentences} - 84.6 \times \frac{syllables}{words}$。输出 0-100（越高越易读） | 与 FKGL 互补 | 同一文本 FRE 和 FKGL 方向相反 |
| FR-403 | **Coleman-Liau Index (CLI)**：$0.0588 \times L - 0.296 \times S - 15.8$（L=每百词字母数，S=每百词句数）。**不需要音节计数** | 作为不依赖音节的补充公式 | CLI 计算结果在合理范围 |
| FR-404 | **音节计数算法**：元音序列计数法（准确率 ~85-90%）。规则：元音 `aeiouy` 连续序列计 1 个音节；词尾 `-e`（非 `-le`）不计；词尾 `-es/-ed` 视上下文处理；最少 1 个音节 | FKGL 和 FRE 依赖音节计数 | "beautiful" → 3 音节；"the" → 1 音节；"smile" → 1 音节 |
| FR-405 | **可读性综合分数**：`readabilityScore = 0.50×fkglToCefr(FKGL) + 0.30×freToCefr(FRE) + 0.20×cliToCefr(CLI)`。多公式互补偿偏差 | 单一公式偏差大 | 三个公式的 CEFR 映射值加权平均 |
| FR-406 | **FKGL 到 CEFR 映射**：Grade 1-2→A1, 3-4→A2, 5-7→B1, 8-10→B2, 11-13→C1, 14+→C2（线性插值） | 学术研究标准映射 | FKGL=6 映射到 CEFR ≈ B1 (3.2) |

### FR-5xx: 长文本采样

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-501 | **自动采样**：文本 > 10,000 词时自动启用采样策略；< 10,000 词全文分析 | 避免长文本评估超时 | 50,000 词文本评估耗时 < 50ms |
| FR-502 | **分层采样**：三段采样（开头 ~1000 词 + 中间 ~1000 词 + 结尾 ~1000 词），覆盖文本整体难度分布 | 开头/中间/结尾难度可能不同 | 采样结果与全文分析结果差异 < 0.5 级别 |
| FR-503 | **采样置信度**：根据采样覆盖率和段间差异计算置信度。段间差异大时（方差 > 1.0）置信度降低到 0.5 以下 | 避免采样误导 | 难度均匀的文本采样置信度 > 0.8；难度差异大的文本置信度 < 0.6 |
| FR-504 | **强制全文分析选项**：提供 `WithFullAnalysis()` option 强制对长文本做全文分析（牺牲性能换精度） | 调用方可在精度优先场景使用 | 50,000 词文本强制全文分析耗时 < 500ms |

### FR-6xx: 文本预处理

| ID | Requirement | Evidence | Testability |
|----|-------------|----------|-------------|
| FR-601 | **分词**：按空格和标点分割。处理连字符词（well-known → 拆分为 well 和 known）、缩写（can't → can not / 跳过）、所有格（'s → 去除） | 基础 NLP | "Well-known author's can't stop..." → ["well", "known", "author", "can", "stop"] |
| FR-602 | **大小写规范化**：所有 token 转小写后查词表 | 词表统一小写存储 | "The" → "the" |

---

## Non-Functional Requirements

| Category | Constraint | Rationale |
|----------|------------|-----------|
| **Performance** | 1,000 词文本评估 < 5ms；10,000 词文本 < 50ms；采样评估恒定 < 20ms | ReadingDays 导入后处理管道中同步使用 |
| **Memory** | 内嵌词表数据（Oxford 5000 + NGSL + AWL + 不规则变体）≤ 2MB 压缩后 | Go binary 大小影响可接受 |
| **Dependencies** | 零第三方依赖（仅用标准库 + `embed`） | 保持库轻量 |
| **Compatibility** | 支持 Go 1.22+（与其他 simp-lee/* 库一致） | 生态一致性 |
| **API Design** | 核心 API ≤ 5 个公开函数/类型；Functional Options 模式 | 与 simp-lee 生态风格一致 |
| **Accuracy** | 在标准 CEFR 分级语料（如 Cambridge English Readers）上评估准确率 > 70%（±1 级别精度）；与人类标注的 Spearman 相关系数 > 0.7 | 学术研究基准 |
| **Testing** | 单元测试覆盖率 > 80%；含 CEFR 标注的参考文本集做 regression 测试 | 算法调参时防止退化 |

---

## Technical Constraints

- **TC1**: 纯 Go 实现，无 CGO 依赖
- **TC2**: 词汇数据使用 Go 1.16+ `embed` 包嵌入。数据格式推荐 CSV 或二进制序列化（gob），启动时解析到 `map[string]Level`
- **TC3**: 不使用 NLP parser，所有句法分析基于启发式规则（关键词匹配、模式匹配）
- **TC4**: Oxford 5000 数据版权归 Oxford University Press，嵌入数据仅包含 `word → CEFR level` 映射（不含释义），学术/个人使用
- **TC5**: NGSL 使用 Creative Commons 许可，可自由使用
- **TC6**: AWL 为学术公开数据，可自由使用
- **TC7**: 分句算法需处理英文常见缩写（内置 ~50 个常见缩写列表：Mr., Mrs., Dr., Jr., Sr., vs., etc., e.g., i.e., U.S., U.K. 等）
- **TC8**: 与 `simp-lee/isdict-api` 的关系：isdict-api 提供义项级别 CEFR 标注（精确到词义），cefr 库提供文本级别评估（宏观难度）。两者互补，不重叠

---

## API 设计参考

```go
// 基础使用
result, err := cefr.Assess("The elaborate plan was discussed by all members, although some disagreed.")

// 带选项
result, err := cefr.Assess(text,
    cefr.WithWeights(0.6, 0.2, 0.2),   // 自定义权重
    cefr.WithFullAnalysis(),             // 强制全文分析
)

// Result 结构体
type Result struct {
    Level      string   // "A1" - "C2"
    Score      float64  // 1.0 - 6.0 连续分数
    Confidence float64  // 0.0 - 1.0

    Vocab       VocabResult       // 词汇维度详情
    Syntax      SyntaxResult      // 句法维度详情
    Readability ReadabilityResult // 可读性维度详情

    WordCount    int  // 总词数
    SentenceCount int // 句数
}

type VocabResult struct {
    Score        float64            // 1.0-6.0
    Distribution map[string]float64 // {"A1": 0.45, "A2": 0.25, ...}
    UnknownRatio float64            // 未知词占比
    ContentWords int                // 内容词数（排除停用词/专有名词/数字）
}

type SyntaxResult struct {
    Score              float64 // 1.0-6.0
    AvgSentenceLength  float64
    SubordinationIndex float64
    PassiveRate        float64
    ConnectorDiversity int
}

type ReadabilityResult struct {
    Score float64 // 1.0-6.0
    FKGL  float64 // Flesch-Kincaid Grade Level
    FRE   float64 // Flesch Reading Ease
    CLI   float64 // Coleman-Liau Index
}
```

---

## Risks & Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| 词汇数据版权（Oxford 5000） | Medium | 仅嵌入 word→level 映射，不含释义；标注数据来源；如有争议可替换为纯 NGSL+AWL+COCA 方案（精度下降但无版权风险） |
| 启发式算法精度不足 | Medium | 输出置信度分数；低置信度时调用方可选择 LLM 辅助评估（ReadingDays 应用层实现） |
| 音节计数不精确 | Low | 启发式方法 85-90% 精确度对统计分析足够；可选嵌入 CMU Pronouncing Dictionary（+5MB）提升精度 |
| 权重参数需要调优 | Medium | 提供可配置权重 API；使用 CEFR 标注语料（Cambridge English Readers）做参数调优 |
| 中文语境中的英文文本特殊性 | Low | 中国英语学习者文本中可能混杂中文；预处理阶段检测并过滤非ASCII字符 |

---

## 词汇数据规格

| 数据源 | 词条数 | CEFR 标注 | 许可 | 嵌入格式 |
|--------|--------|-----------|------|---------|
| Oxford 5000 | ~5,000 | A1-C1（精确标注） | OUP 版权（仅嵌入 word→level 映射） | CSV → `map[string]int` |
| NGSL | ~2,800 | 无（按频率排名映射）| Creative Commons，免费 | CSV → `map[string]int` |
| AWL | ~570 词族（~3,000 词形）| 无（按 sublist 映射 B2-C1）| 学术公开 | CSV → `map[string]int` |
| 不规则变体表 | ~1,000 | 无（映射到 lemma 后查表）| 自建 | CSV → `map[string]string` |
| 停用词表 | ~200 | 无（标记跳过）| 公共领域 | `map[string]bool` |
| 常见缩写表 | ~50 | 无（分句时识别）| 自建 | `map[string]bool` |
| 从属连词/关系代词 | ~30 | 无（句法分析标记词）| 自建 | `map[string]bool` |
| be 动词形式 | ~10 | 无（被动语态检测）| 自建 | `map[string]bool` |
| 不规则过去分词 | ~200 | 无（被动语态检测）| 自建 | `map[string]bool` |

---

## 研究资料

| 来源 | 关键贡献 |
|------|---------|
| Vajjala & Meurers (2012) | 多特征 CEFR 分类方法论（SVM + 词汇/句法/可读性特征），~83% 准确率 |
| Pilán, Volodina & Zesch (2016) | 词汇覆盖率特征的跨语言有效性验证 |
| Xia, Kochmar & Briscoe (2016) | 特征重要性排名：词汇 > 句法 > 话语 |
| Coxhead (2000) | AWL 原始论文 |
| Browne, Culligan & Phillips (2013) | NGSL 发布与方法论 |
| Kincaid et al. (1975) | Flesch-Kincaid 公式原始论文 |
| Council of Europe (2001, 2020) | CEFR 框架官方文档 |
| `.agents-work/research/cefr-assessment-algorithms.md` | CEFR 评估算法完整技术研究报告 |
