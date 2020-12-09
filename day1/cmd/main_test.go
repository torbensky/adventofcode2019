package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	got := part1(strings.NewReader(testData))
	want := 3434390
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.NewReader(testData))
	want := 5148724
	if got != want {
		t.Fatalf("expected %d got %d\n", want, got)
	}
}

const testData = `98541
129056
134974
66390
121382
94570
107586
98767
65101
56320
63431
112200
119262
142745
143941
148764
70301
149623
125170
114562
136701
76971
52292
127671
107547
77460
55268
119986
104257
82814
64527
74279
98542
54710
96317
105670
146248
134587
104028
65286
91788
106723
137825
139949
74403
106574
133990
96165
121316
94072
76612
109470
147556
113157
67117
85237
134232
94622
76160
107532
120637
51505
82847
105600
97719
113114
68177
149213
116125
145577
83921
134810
138804
90125
70621
103245
51584
93437
125352
100578
53497
112023
92999
107998
148030
101185
65777
92272
145846
81488
61957
69551
125625
146328
123666
102629
121996
94172
128023
123472`
