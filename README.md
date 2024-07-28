gotest-output-parser
====================

Takes a `go test` .json output file, and dumps the compact test log - printing successes and logs only for failed tests.

```sh
# Run your tests, generate JSON file.
$ go test ./... -json > testoutput.json

# Parse the JSON file for a compact dump.
$ go build && cat ../testoutput.json | ./testparser

PASS: TestBitstring (0.00s)
PASS: TestMerkleTreeAccumulate (0.00s)
PASS: TestCreateWallet (0.00s)
PASS: TestSign (0.00s)
PASS: TestVerifyWithRealSig (0.00s)
PASS: TestVerify (0.00s)
PASS: TestRecreateWallet (0.00s)
PASS: TestOpenDB (0.00s)
PASS: TestDagLatestTipIsSet (0.00s)
FAIL: TestDagAddBlockUnknownParent (0.00s)

=== RUN   TestDagAddBlockUnknownParent
2024/07/28 18:58:12 [blockdag] (db) Database version: 0
2024/07/28 18:58:12 [blockdag] (db) Running migration: 1
2024/07/28 18:58:12 [blockdag] (db) Database upgraded to: 1
2024/07/28 18:58:12 [blockdag] (db) Database upgraded to: 2
2024/07/28 18:58:12 [pow] SolvePOW target: 7237005577332262213973186563042994240829374041602535252466099000494570602495
2024/07/28 18:58:12 [pow] Solved in 21 iterations
2024/07/28 18:58:12 [pow] Hash: 33383330323733393533303235393234323530343339353739383634303631353031373136313632363033333037313231323238363536323331333839373536383433383934373832313531
2024/07/28 18:58:12 [pow] Nonce: 21
2024/07/28 18:58:12 [pow] VerifyPOW target: 7237005577332262213973186563042994240829374041602535252466099000494570602495
Genesis block hash=0877dbb50dc6df9056f4caf55f698d5451a38015f8e536e9c82ca3f5265c38c7 work=30
2024/07/28 18:58:12 [pow] SolvePOW target: 7237005577332262213973186563042994240829374041602535252466099000494570602495
2024/07/28 18:58:12 [pow] Solved in 21 iterations
2024/07/28 18:58:12 [pow] Hash: 33383330323733393533303235393234323530343339353739383634303631353031373136313632363033333037313231323238363536323331333839373536383433383934373832313531
2024/07/28 18:58:12 [pow] Nonce: 21
2024/07/28 18:58:12 [pow] VerifyPOW target: 7237005577332262213973186563042994240829374041602535252466099000494570602495
Genesis block hash=0877dbb50dc6df9056f4caf55f698d5451a38015f8e536e9c82ca3f5265c38c7 work=30
2024/07/28 18:58:12 [blockdag] Initialising block DAG...
2024/07/28 18:58:12 [blockdag] Inserted genesis epoch difficulty=7237005577332262213973186563042994240829374041602535252466099000494570602495
2024/07/28 18:58:12 [blockdag] Inserted genesis block hash=0877dbb50dc6df9056f4caf55f698d5451a38015f8e536e9c82ca3f5265c38c7 work=30
2024/07/28 18:58:12 [blockdag] New headers tip: height=0 hash=0877dbb50dc6df9056f4caf55f698d5451a38015f8e536e9c82ca3f5265c38c7
2024/07/28 18:58:12 [blockdag] New full tip: height=0 hash=0877dbb50dc6df9056f4caf55f698d5451a38015f8e536e9c82ca3f5265c38c7
    blockdag_test.go:216:
        	Error Trace:	/Users/liamz/tinychain-go/core/nakamoto/blockdag_test.go:216
        	Error:      	Not equal:
        	            	expected: "Unknown parent block."
        	            	actual  : "Block not found."

        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -1 +1 @@
        	            	-Unknown parent block.
        	            	+Block not found.
        	Test:       	TestDagAddBlockUnknownParent
--- FAIL: TestDagAddBlockUnknownParent (0.00s)

PASS: TestDagAddBlockTxCount (0.00s)
PASS: TestDagAddBlockTxsValid (0.00s)
PASS: TestDagAddBlockTxMerkleRootValid (0.00s)
PASS: TestDagAddBlockSuccess (0.00s)
PASS: TestDagAddBlockWithDynamicSignature (0.00s)
PASS: TestDagGetBlockByHashGenesis (0.00s)
PASS: TestDagBlockDAGInitialised (0.00s)
PASS: TestDagGetEpochForBlockHashGenesis (0.00s)
PASS: TestDagGetEpochForBlockHashNewBlock (0.00s)
```