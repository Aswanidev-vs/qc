# QuickCalc — Work Plan

## Phase 1: Codebase audit & prod-readiness baseline (agreed)
- [x] Identify prod-readiness issues in errors/API consistency, numerical tolerances, numerical robustness, and testing gaps
- [x] Introduce a single source of truth for numerical tolerances (constants)
- [ ] Add/complete GoDoc for exported functions (where missing/incomplete)
- [ ] Add edge-case/invalid-input/error-path tests without changing algorithms
- [ ] Normalize error returns to use exported sentinel errors where feasible (without altering math behavior)

## Phase 2: Prod-ready engineering upgrades (after Phase 1 approval)
- [ ] Replace matrix determinant cofactor expansion with LU/Gaussian elimination (O(n^3))
- [ ] Standardize failure modes/pivoting tolerances across matrix inverse/determinant/rank
- [ ] Refactor heavy calculus paths for reduced recursion/allocations (only after Phase 2 approval)

## Phase 3: Add more mathematical content (after Phase 2 approval)
- [ ] Add linear algebra upgrades (least squares, Gram-Schmidt, 2x2 eigenvalues, linear solves)
- [ ] Add calculus upgrades (multivariate Jacobian/gradient, adaptive ODE integration)
- [ ] Add probability/statistics upgrades (extra distributions + at least one hypothesis test)

## Phase 4: Hardening (after Phase 3 approval)
- [ ] Expand test coverage (boundary + numerical regression + failure cases)
- [ ] Add CI checks (gofmt, go vet, go test ./..., coverage thresholds)
- [ ] Define and apply “prod-ready” success criteria
