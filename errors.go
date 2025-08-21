package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// 가상의 원본 에러 (database/sql에서 올 수 있는 에러)
var originalError = errors.New("connection refused")

// errors.As 테스트 - 커스텀 에러 타입
type CustomError struct {
	Code int
	Msg  string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("code %d: %s", c.Code, c.Msg)
}

func demonstrateErrorComparison() {
	fmt.Println("=== 에러 출력 방식 비교 ===\n")

	// 방법 1: errors.Wrap 사용
	err1 := errors.Wrap(originalError, "failed to query expired notification requests")

	// 방법 2: errors.WithMessage + errors.WithStack 조합
	err2 := errors.WithMessage(errors.WithStack(originalError), "failed to query expired notification requests")

	// 기본 에러 메시지 출력 비교
	fmt.Println("📝 기본 에러 메시지:")
	fmt.Printf("Wrap:                 %v\n", err1)
	fmt.Printf("WithMessage+WithStack: %v\n", err2)
	fmt.Println()

	// Error() 메서드 출력 비교
	fmt.Println("📝 Error() 메서드:")
	fmt.Printf("Wrap:                 %s\n", err1.Error())
	fmt.Printf("WithMessage+WithStack: %s\n", err2.Error())
	fmt.Println()

	// 상세 스택트레이스 출력 비교 (%+v)
	fmt.Println("📝 스택트레이스 포함 출력:")
	fmt.Println("--- errors.Wrap ---")
	fmt.Printf("%+v\n\n", err1)

	fmt.Println("--- errors.WithMessage + errors.WithStack ---")
	fmt.Printf("%+v\n\n", err2)

	// 에러 타입 정보 비교
	fmt.Println("📝 에러 구조 정보:")
	fmt.Printf("Wrap type:                 %T\n", err1)
	fmt.Printf("WithMessage+WithStack type: %T\n", err2)
	fmt.Println()

	// Cause 추적 비교
	fmt.Println("📝 원본 에러 추적:")
	fmt.Printf("Wrap cause:                 %v\n", errors.Cause(err1))
	fmt.Printf("WithMessage+WithStack cause: %v\n", errors.Cause(err2))
	fmt.Println()
}

func demonstrateNestedErrors() {
	fmt.Println("=== 중첩된 에러 시나리오 ===\n")

	// 여러 레벨에서 에러 전파 시뮬레이션
	level3Err := errors.New("database connection timeout")

	// 방법 1: errors.Wrap 체인
	level2Err1 := errors.Wrap(level3Err, "repository layer failed")
	level1Err1 := errors.Wrap(level2Err1, "service layer failed")

	// 방법 2: WithMessage + WithStack 체인
	level2Err2 := errors.WithMessage(errors.WithStack(level3Err), "repository layer failed")
	level1Err2 := errors.WithMessage(errors.WithStack(level2Err2), "service layer failed")

	fmt.Println("📝 중첩 에러 메시지:")
	fmt.Printf("Wrap chain:                 %v\n", level1Err1)
	fmt.Printf("WithMessage+WithStack chain: %v\n", level1Err2)
	fmt.Println()

	fmt.Println("📝 중첩 에러 스택트레이스:")
	fmt.Println("--- Wrap chain ---")
	fmt.Printf("%+v\n\n", level1Err1)

	fmt.Println("--- WithMessage+WithStack chain ---")
	fmt.Printf("%+v\n\n", level1Err2)
}

func demonstrateStackDuplication() {
	fmt.Println("=== 중복 스택 방지 테스트 ===\n")

	// 이미 스택이 있는 에러
	stackedErr := errors.WithStack(originalError)

	// 방법 1: errors.Wrap (중복 스택 방지됨)
	err1 := errors.Wrap(stackedErr, "additional context")

	// 방법 2: WithMessage + WithStack (중복 스택 방지됨)
	err2 := errors.WithMessage(errors.WithStack(stackedErr), "additional context")

	fmt.Println("📝 중복 스택 방지 확인:")
	fmt.Printf("Original stacked error: %T\n", stackedErr)
	fmt.Printf("Wrap on stacked:        %T\n", err1)
	fmt.Printf("WithMessage+WithStack:  %T\n", err2)
	fmt.Println()

	fmt.Println("📝 에러 메시지:")
	fmt.Printf("Wrap:                 %v\n", err1)
	fmt.Printf("WithMessage+WithStack: %v\n", err2)
	fmt.Println()
}

func demonstratePerformance() {
	fmt.Println("=== 성능 비교 (간단한 측정) ===\n")

	iterations := 10000

	// Wrap 방식 측정
	fmt.Printf("errors.Wrap 방식 (%d회):\n", iterations)
	for i := 0; i < iterations; i++ {
		_ = errors.Wrap(originalError, "test message")
	}
	fmt.Println("완료")

	// WithMessage + WithStack 방식 측정
	fmt.Printf("WithMessage+WithStack 방식 (%d회):\n", iterations)
	for i := 0; i < iterations; i++ {
		_ = errors.WithMessage(errors.WithStack(originalError), "test message")
	}
	fmt.Println("완료")

	fmt.Println("(실제 성능 차이는 미미함)")
	fmt.Println()
}

func demonstrateErrorIs() {
	fmt.Println("=== errors.Is/As 동작 비교 ===\n")

	var sentinelError = errors.New("sentinel error")

	// 방법 1: errors.Wrap
	err1 := errors.Wrap(sentinelError, "wrapped with Wrap")

	// 방법 2: WithMessage + WithStack
	err2 := errors.WithMessage(errors.WithStack(sentinelError), "wrapped with WithMessage+WithStack")

	fmt.Println("📝 errors.Is 동작:")
	fmt.Printf("errors.Is(err1, sentinelError): %v\n", errors.Is(err1, sentinelError))
	fmt.Printf("errors.Is(err2, sentinelError): %v\n", errors.Is(err2, sentinelError))
	fmt.Println()

	customErr := CustomError{Code: 500, Msg: "internal error"}
	err3 := errors.Wrap(customErr, "wrapped custom error")
	err4 := errors.WithMessage(errors.WithStack(customErr), "wrapped custom error")

	var target CustomError
	fmt.Println("📝 errors.As 동작:")
	fmt.Printf("errors.As(err3, &target): %v\n", errors.As(err3, &target))
	fmt.Printf("errors.As(err4, &target): %v\n", errors.As(err4, &target))
}

func main() {
	demonstrateErrorComparison()
	demonstrateNestedErrors()
	demonstrateStackDuplication()
	demonstratePerformance()
	demonstrateErrorIs()
}
