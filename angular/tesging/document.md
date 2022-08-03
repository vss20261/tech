# Docs

___
## 기본 정의
 * web front 테스트 유형
   * 단위 테스트 : 컴포넌트, 함수 단위의 작은 코드
   * e2e 테스트 : 사용자의 행동에 따라 애플리케이션의 각 부분이 정해진 대로 동작하는지 확인
   * 부하 테스트
* 테스트 프레임워크 : **Jasmine**
    - 자바 스크립트 테스트 프레임워크
    - angular 가이드 문서에서 Jasmine을 권장
* 테스트 러너 : **Karma**
    - 원래는 AngularJS 팀 내부에서 사용하려고 만든 것이지만 보편화 됨.
    - Node.js에서 구동되며, 자바 스크립트 테스트 프레임워크를 지원한다.
    - browser 관경에서도 자유롭고 desktop, phone, tablet에서도 수행이 가능함.
    - 여러 브라우저를 띄워서 동작 확인 가능.
    [ 동작 순서 ]
    1. karma 전용 서버 띄움
    2. 테스트 소스코드를 karma.config.js에 미리 정의한 웹 브라우저 IFrame 내부로 불러들여 테스트를 실행한다.
    3. 테스트 수행 후, 수행 결과를 karma 서버에서 받아 콘솔을 통해 보여줌.
* 라이브러리 : **Spectator**
     - Spectator 는 Angular 애플리케이션을 테스트하기 위한 라이브러리
     - 간소화된 단위 테스트를 수행
     - Jest, Jasmine Support

---
## 목차

* [테스트 시 사용되는 메소드 정리](#테스트-시-사용되는-메소드-정리)
* [자주쓰는 command](#자주쓰는-command)
* [테스트 범위 지정](#테스트-범위-지정)
* [partial](#partial)
* [jasmine 예약어](#jasmine-예약어)
* [matcher](#matcher)
  * [toEqual](#toequal)
  * [toBe](#tobe)
  * [toBeTruthy](#tobetruthy)
  * [부정 not](#부정-not)
  * [toContain](#tocontain)
  * [toHaveSize](#tohavesize)
  * [toBeDefined](#tobedefined)
  * [toBeNull](#tobenull)
  * [toBeNaN](#tobenan)
  * [toBeGreaterThan](#tobegreaterthan)
  * [toBeLessThan](#tobelessthan)
  * [toBeCloseTo](#tobecloseto)
  * [toMatch](#tomatch)
  * [toThrow](#tothrow)
  * [toHaveBeenCalled](#tohavebeencalled)
  * [toHaveBeenCalledTimes](#tohavebeencalledtimes)
* [reference](#reference)

---
### 테스트 시 사용되는 메소드 정리
```
* 테스트 실행
describe()

* 테스트 스펙(테스트 케이스) 정의
it()

* 테스트 검증
expect(1).toEqual(1)

* 테스트 실행
    * 테스트 스윗(suite; 하나 이상의 테스트 스펙을 모은 것)에 해당.
    * 이 함수 안에 테스트를 실행하는 코드를 작성.
* 테스트 스펙 정의
    * 개별 테스트 케이스를 작성한다.
    * 코드 실행 결과를 검증하는 내용을 기술한다.
* 테스트 검증
    * 검증에 실패하면 프레임워크에서 에러를 발생시킨다.
    * expect() 함수 뒤에 매처(matcher)를 붙여 사용한다.

// 테스트 셋업 함수
* 테스트 스윗에 실행되기 전 1번만 실행 됨
beforeAll()

* 테스트 스펙이 실행되기 전에 1번씩 실행 됨
beforeEach()

* 테스트 스윗이 끝나고 1번만 실행 됨
afterAll()

* 테스트 스펙이 끝나기 전에 1번씩 실행 됨
afterEach()
```

---
### 자주쓰는 command

~~~
 * test 실행
  $ ng test box

 * code coverage 확인
  $ ng test box --code-coverage

 * browser없이 실행
  $ ng test box --watch=false
~~~

---
### 테스트 범위 지정
```typescript
*** 특정범위 테스트 ***
// angular.json (project = box)
 "test": {
          "builder": "@angular-devkit/build-angular:karma",
          "options": {
            // 여기부터
            "include": [
              "projects/box/src/app/psnz/download-box/**/*.spec.ts" // 테스트를 진행하고 싶은 특정 경로
            ],
            // 여기까지
            "codeCoverageExclude": [
              "projects/box/src/app/_shared/**/*", // 코드 커버리지에 들어가는게 싫다 하면 넣기
              "projects/box/src/app/_core/**/*", // 코드 커버리지에 들어가는게 싫다 하면 넣기
              "projects/box/src/app/**/*-api.service.ts",
              "projects/box/src/app/**/*list.service.ts"
            ],
```

```typescript
*** 단일 테스트 케이스 ***
이를 수행하는 간단한 방법은 설명 을 시작 하고 모든 것을 f 로 시작

fdescribe('testecomponente', () => {
  fit('should create', () => {
    //your code 
  }
  
}
```

```typescript
*** skip 테스트 ***
이를 수행하는 간단한 방법은 설명 을 시작 하고 모든 것을 x 로 시작

xdescribe('testecomponente', () => {
  xit('should create', () => {
    //your code 
  }
  
}
```

---
## partial
```typescript
TypeScript는 일반적인 유형 변환을 용이하게 하기 위해 여러 유틸리티 타입을 제공
유틸리티 타입중 partial의 대한 예제코드

type Partial<T> = {
  [A in keyof T]?: T[A];
}

interface Member {
  id: number;
  name: string;
}

export const MockInterface = <T>(p?: Partial<T>): T => {
  return <T>{ ...p };
};

// Member interface의 구현된 모든 필드를 입력할 필요가 없어진다.
const grooo = MockInterface<Member>({name: grooo});

// partial을 이용하면 다음과 같은 interface를 구현하게 된다고 생각하면 된다.
interface Member {
  id?: number;
  name?: string;
}
```

---
### jasmine 예약어
~~~
Jasmine 예약어
다음은 Jasmine과 충돌을 일으키지 않도록 코드에서 사용해서는 안되는 단어입니다.

jasmine(및 네임스페이스의 모든 것)
describe
it
expect
beforeEach
afterEach
runs
waits
waitsFor
spyOn
xdescribe
xit
~~~

---
## matcher

### toEqual
```typescript
expect(true).toEqual(true); // success
expect([1, 2, 3]).toEqual([1, 2, 3]); // success
expect([1, 2, 3, 4]).toEqual([1, 2, 3]); // failure
```

### toBe
```typescript
const spot = { species: "Border Collie" };
const cosmo = { species: "Border Collie" };

expect(spot).toBe(cosmo);     // failure not the same object
expect(spot).toBe(spot);      // success the same object

const arr = [1, 2, 3];
expect(arr).toEqual([1, 2, 3]);  // success; equivalent
expect(arr).toBe([1, 2, 3]);     // failure; not the same array
```

### toBeTruthy
```typescript
expect(true).toBeTruthy(); // success
expect(12).toBeTruthy(); // success
expect({}).toBeTruthy(); // success

expect("").toBeTruthy(); // failure
expect(0).toBeTruthy(); // failure
expect(null).toBeTruthy(); // failure
expect(undefined).toBeTruthy(); // failure
```
### 부정 not
```typescript
expect("hi").not.toContain("hello"); // success
```
### toContain

```typescript
//array
expect([1, 2, 3, 4]).toContain(3);
expect(["Penguin", "Turtle", "Pig", "Duck"]).toContain("Duck"); // success

// object
const dog = { name: "Fido" };
expect([
    { name: "Spike" },
    { name: "Fido" },
    { name: "Spot" }
]).toContain(dog); // success

// string
expect("Hello world").toContain("world");
```

### toHaveSize
```typescript
const array = [1,2,3];
const object = {a: 1, b:2, c: {d: 3}}

expect(array).toHaveSize(3) // success
expect(object).toHaveSize(3) // success

```
### toBeDefined
```typescript
const somethingUndefined;
expect("Hello!").toBeDefined();                  // success
expect(null).toBeDefined();                      // success
expect(somethingUndefined).toBeDefined();        // failure
```

### toBeNull
```typescript
expect(null).toBeNull();                // success
expect(false).toBeNull();               // failure
expect(somethingUndefined).toBeNull();  // failure
```

### toBeNaN
```typescript
expect(5).not.toBeNaN();              // success
expect(0 / 0).toBeNaN();              // success
expect(parseInt("hello")).toBeNaN();  // success
```
### toBeGreaterThan
```typescript
expect(8).toBeGreaterThan(6); // success
expect(8).toBeGreaterThan(8); // failure
```
### toBeLessThan
```typescript
expect(5).toBeGreaterThan(12); // success
expect(8).toBeGreaterThan(8); // failure
```
### toBeCloseTo
```typescript
expect(12.34).toBeCloseTo(12.32, 1);  // success
expect(12.3456789).toBeCloseTo(12, 0);  // success
expect(12.34).toBeCloseTo(12.3, 3);  // failure
expect(12.34).toBeCloseTo(12.3, 4);  // failure
```
### toMatch
```typescript
expect("foo bar").toMatch(/bar/); // success
expect("jasmine@example.com").toMatch("\w+@\w+\.\w+"); // success
```
### toThrow
```typescript
const throwMeAnError = function() {
    throw new Error();
};
expect(throwMeAnError).toThrow(); // success
```
### toHaveBeenCalled
```typescript
// 호출이 되었는지 확인
expect(spectator.component.function).toHaveBeenCalled()
```
### toHaveBeenCalledTimes
```typescript
// 지정된 횟수만큼 호출이 되었는지 확인
expect(spectator.component.function).toHaveBeenCalledTimes(2)
```

---
## reference
* jasmine 
  * [Docs](https://jasmine.github.io/pages/docs_home.html)
  * [Guide](https://testing-angular.com/introduction/)
* spectator
  * [Docs](https://ngneat.github.io/spectator/docs/installation)
  * [Github](https://github.com/ngneat/spectator)
  * [DetectChanges](spectator/detectChanges.md)
* partial
  * [Partial](https://www.becomebetterprogrammer.com/typescript-partial-type/) 
  * [Utility Types](https://www.typescriptlang.org/docs/handbook/utility-types.html)
* angular
  * [Testing components](https://angular.kr/guide/testing-components-scenarios)
  * [Command](https://angular.io/cli/test)
