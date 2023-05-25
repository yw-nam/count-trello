# [Trello] 특정 라벨의 목록별 카드 개수 세기 프로그램

## 필요한 환경변수

- TOKEN : trello api token
- API_KEY : trello api key
- BOARD_ID : 해당 trello board의 ID
- BASE_DATE : 생성일로 부터 지난일자 계산 기준일 (fmt:YYYYMMDD, def:실행일자)
- MAX_WEEK : 카우트할 최대 주 수 (def:-1 --> 정렬없이 전부출력)

## 사용법

    TOKEN={트랠로 토큰 값} \
    API_KEY={트랠로 API 키 값} \
    BOARD_ID={트랠로 보드 ID} \
    LABEL={보려는 라벨} \
    go run main.go
