package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"poten-invitation-golang/app/expense/controller"
	"poten-invitation-golang/app/expense/repository"
	"poten-invitation-golang/app/expense/service"
	"poten-invitation-golang/app/external"
	"poten-invitation-golang/util"
)

func main() {
	// 현재 디렉토리 경로 가져오기
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("현재 디렉토리 경로 가져오기 실패:", err)
		return
	}

	// 프로젝트 폴더 구조를 저장할 슬라이스
	projectStructure := make([]string, 0)

	// 현재 디렉토리부터 하위 디렉토리들을 순회하며 폴더 구조를 뽑아냅니다.
	filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		// 에러 처리
		if err != nil {
			fmt.Println("폴더 순회 중 에러 발생:", err)
			return err
		}

		// 파일이 아닌 폴더만 처리
		if !info.IsDir() {
			return nil
		}

		// 현재 디렉토리 경로를 제외하고 출력
		relPath, err := filepath.Rel(currentDir, path)
		if err != nil {
			fmt.Println("상대 경로 변환 실패:", err)
			return err
		}

		// 프로젝트 폴더 구조에 추가
		projectStructure = append(projectStructure, relPath)
		return nil
	})

	// 프로젝트 폴더 구조 출력
	fmt.Println("프로젝트 폴더 구조:")
	for _, path := range projectStructure {
		fmt.Println(path)
	}

	if err := util.EnvInitializer(); err != nil {
		panic(err)
	}
	db := external.NewDB()
	expenseRepository := repository.NewExpenseRepository(db)
	expenseService := service.NewExpenseService(expenseRepository)
	expenseController := controller.NewExpenseController(expenseService)

	if err := external.GetRouter(expenseController).Run(); err != nil {
		log.Fatalln(err)
	}

}
