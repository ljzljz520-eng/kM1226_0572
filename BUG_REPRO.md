# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
--- FAIL: TestWeaponUpgradeTouchesOnlySelectedSlot (0.01s)
    weapon_regression_test.go:37: adjacent weapon changed from 12 to 20
FAIL
FAIL	spacetrash	0.009s
?   	spacetrash/cmd/spacedock	[no test files]
?   	spacetrash/rules	[no test files]
ok  	spacetrash/game	0.001s
ok  	spacetrash/input	0.001s
ok  	spacetrash/model	0.001s
ok  	spacetrash/report	0.006s
ok  	spacetrash/service	0.016s
ok  	spacetrash/storage	0.008s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/spacedock): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/spacedock): exit `0`
