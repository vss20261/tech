

### Bad
~~~
  describe("계산기 추가", () => {
      it("양의 정수를 더하고, 빼고, 곱하고, 나눌 수 있습니다.", () => {
          let calc = new Calculator;
          expect(calc.add(2, 3)).toEqual(5);
          expect(calc.sub(8, 5)).toEqual(3);
          expect(calc.mult(4, 3)).toEqual(12);
          expect(calc.div(12, 4)).toEqual(3);
      });
  });
~~~

### Good
~~~
  describe("계산기 추가", function() {
      let calc;
      beforeEach(() => {
          calc = new Calculator();
      });
      it("양의 정수를 추가할 수 있습니다.", () => {
          expect(calc.add(2, 3)).toEqual(5);
      });
      it("양의 정수를 뺄 수 있습니다.", () => {
          expect(calc.sub(8, 5)).toEqual(3);
      });
      it("양의 정수를 곱할 수 있습니다.", () => {
          expect(calc.mult(4, 3)).toEqual(12);
      });
      it("양의 정수를 나눌 수 있습니다", () => {
          expect(calc.div(12, 4)).toEqual(3);
      });
  });
~~~