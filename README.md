# [Trello] 특정 라벨의 목록별 카드 개수 세기 프로그램

## 필요한 환경변수

- TOKEN : trello api token
- API_KEY : trello api key
- BOARD_ID : 해당 trello board의 ID

- BASE_DATE : 생성일로 부터 지난일자 계산 기준일 (fmt:YYYYMMDD, def:실행일자+1, 해당일자 00:00 기준)
- MAX_WEEK : 카우트할 최대 주 수 (def:-1 --> 정렬없이 전부출력)
- REQ_SPEED : 10초당 최대 요청 수 (def:90 , max:100)

## 사용법

    TOKEN={트랠로 토큰 값} \
    API_KEY={트랠로 API 키 값} \
    BOARD_ID={트랠로 보드 ID} \
    LABEL={보려는 라벨} \
    MAX_WEEK=8 \
    go run main.go

## TODO

- 2차 기능의 날짜 정합성 검증용 테스트 작성
