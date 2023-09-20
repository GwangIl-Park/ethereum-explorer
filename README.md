# ethereum-explorer

- etherscan의 화면을 기반으로 작업할 예정  https://etherscan.io/

- 메인페이지
    - 상위 6개 블록 확인 가능
        - blocknumber 클릭시 블록 정보로 이동
        - account 클릭시 account 정보로 이동
        - 전체 보기 시 블록리스트 정보로 이동
    - 상위 6개 트랜잭션 확인 가능
        - txHash 클릭시 트랜잭션 정보로 이동
        - account 클릭시 account 정보로 이동
        - 전체 보기 시 트랜잭션 리스트 정보로 이동
- 블록
    - 기본적인 블록정보 확인 가능
    - 블록의 트랜잭션 개수 확인 가능
        - 클릭 시 특정 블록 트랜잭션 리스트로 이동
    - 블록의 internal 트랜잭션 개수
        - 클릭 시 특정 블록 internal 트랜잭션 리스트로 이동
    - receipient 주소 확인 가능
        - 클릭 시 account 정보로 이동
- 블록 리스트
    - 블록 리스트 확인 가능
    - 블록넘버 클릭시 블록 정보로 이동
    - tx 개수 클릭시 특정 블록 트랜잭션 리스트로 이동
    - account 클릭시 account 정보로 이동
- 트랜잭션
    - 기본적인 트랜잭션 정보 확인 가능
    - 블록넘버 확인 가능
        - 클릭 시 블록 정보로 이동
    - from account 정보 확인 가능
        - 클릭시 account 정보로 이동
    - to account 정보 확인 가능
        - 클릭시 account 정보로 이동
    - 로그 정보 확인 가능
    - internal 트랜잭션 확인 가능
- 트랜잭션 리스트
    
    
- account

### api

|Name|URL|Method|
|---|---|---|
|메인 페이지|/|GET|
|모든 블록 리스트|/blocks|GET|
|특정 블록 정보|/block/{blockNumber}|GET|
|특정 블록 internal transaction|/internaltxs?block={blockNumber}|GET|
|모든 트랜잭션 리스트|/txs|GET|
|특정 트랜잭션 정보|/tx/{hash}|GET|
|특정 블록 트랜잭션 리스트|/txs?block={blockNumber}|GET|
|account 정보|/address/{address}|GET|

### DB 설계

![image](https://github.com/GwangIl-Park/ethereum-explorer/assets/40749130/41c0921a-6729-4383-9801-c5f9bb2d3484)
