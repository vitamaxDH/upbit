# Max's upbit api

This project was created merely to have fun with upbit's apis in Go.
I do not take any responsibility of any asset loss using apis in this repository.


## Strategy

1. Fetch the 10-minute candles every 10 minutes
2. Check the 3 latest candles
3. 최신 봉이 앞 2개 봉의 평균과 비슷할 경우 해당 마켓의 평균가 업데이트.
4. 최신 봉이 앞 2개 봉의 평균보다 갑자기 높을 경우 (퍼센트는 추후 결정)
    - 평균가 업데이트 flag column to false
    - 마켓 테이블에 monitoring_candle_count column 1 증가
    - 20분 더 기다린 후 기록해두었던 평균가보다 특정 퍼센트 이상 높을 경우 매매

## Model

1. 구매 - default 3시간 보유 후 판매

| column                  | type      | desc               |
|:------------------------|:----------|:-------------------|
| id                      | number    | 아이디              |
| market_id               | number    | 마켓 아이디         |
| avg_price               | decimal   | 평균가 업데이트 여부 |
| updates_avg_price       | boolean   | 평균가 업데이트 여부 |
| monitoring_candle_count | number    | 아이디              |
| created_at              | timestamp | 생성 시간           |
| updated_at              | timestamp | 수정 시간           |

2. 마켓 - market 정보 (market_model.go 와 동일)