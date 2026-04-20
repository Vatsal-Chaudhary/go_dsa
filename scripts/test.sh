#!/usr/bin/env bash
set -euo pipefail

file_path="${1:-}"
func_name="${2:-}"

if [[ -z "$file_path" ]]; then
	printf 'Usage: make test FILE=path/to/file_test.go [FUNC=TestName]\n' >&2
	printf '   or: make test-file FILE=path/to/file_test.go\n' >&2
	printf '   or: make test-func FILE=path/to/file_test.go FUNC=TestName\n' >&2
	exit 1
fi

if [[ ! -f "$file_path" ]]; then
	printf 'Test file not found: %s\n' "$file_path" >&2
	exit 1
fi

if [[ "$file_path" != *_test.go ]]; then
	printf 'Expected a Go test file ending in _test.go: %s\n' "$file_path" >&2
	exit 1
fi

pkg_dir=$(dirname "$file_path")

if [[ -n "$func_name" ]]; then
	printf 'Running %s from %s\n' "$func_name" "$file_path"
	exec go test "./$pkg_dir" -run "^${func_name}$"
fi

mapfile -t test_names < <(sed -nE 's/^func (Test[^ (]*).*/\1/p' "$file_path")

if [[ ${#test_names[@]} -eq 0 ]]; then
	printf 'No top-level Test* functions found in %s\n' "$file_path" >&2
	exit 1
fi

run_pattern="^($(IFS='|'; printf '%s' "${test_names[*]}"))$"

printf 'Running tests from %s\n' "$file_path"
	exec go test "./$pkg_dir" -run "$run_pattern"
