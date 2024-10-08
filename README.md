# arrow_copy  
[다운로드](https://github.com/jinwoole/arrow_copy/releases/latest/download/arrowcopy_1.0.0.exe)

arrow_copy는 간단하고 효율적인 Go 애플리케이션으로, 자주 사용되는 유니코드 화살표 문자를 쉽게 복사할 수 있게 해줍니다. 작고 깔끔한 GUI를 제공하여 데스크톱에서 편리하게 사용할 수 있습니다.  
![image](https://github.com/user-attachments/assets/ec356a8f-f04c-40c7-80dc-34c18f3de681)

## 기능
- 4개의 기본 방향 화살표(↑, ←, →, ↓) 제공
- 클릭 한 번으로 클립보드에 화살표 복사  
- 간편한 창 이동을 위한 드래그 핸들  

## 설치 요구사항
- Windows 10, 11  

## 혹시 빌드하고 싶다면?

1. 이 저장소를 클론합니다
2. 설정파일에 대충 써 있는 의존성을 설치합니다
3. 리소스 파일을 생성합니다
   ```
   rsrc -manifest arrow-copier.manifest -o rsrc.syso
   ```
4. 프로그램을 빌드합니다, 옵션 안붙이면 터미널 나와요:
   ```
   go build -ldflags="-H windowsgui"
   ```

## 사용 방법
1. 실행합니다.
2. 작은 창이 나타나면, 원하는 화살표 버튼을 클릭합니다.
3. 클릭한 화살표가 클립보드에 복사됩니다.
4. 창 하단의 "=" 표시 버튼을 드래그하여 창을 손쉽게 이동할 수 있습니다.

## 기여
이유는 잘 모르겠지만, 그래도 버그 리포트, 기능 제안 또는 풀 리퀘스트는 언제나 환영합니다.
