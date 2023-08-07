#include <stdio.h>

double calc(char op, double num1, double num2)
{
  double result;
  switch (op)
  {
  case '+':
    result = num1 + num2;
    break;
  case '-':
    result = num1 - num2;
    break;
  case '*':
    result = num1 * num2;
    break;
  case '/':
    result = num1 / num2;
    break;
  default:
    result = 0;
    break;
  }
  return result;
}

int main() {

  char op;
  double num1;
  double num2;

  printf("what is your op?\n");
  scanf("%c", &op);

  printf("\nthe first number:\n");
  scanf("%lf", &num1);

  printf("\nthe second number:\n");
  scanf("%lf", &num2);

  double result = calc(op, num1, num2);

  printf("\nthe result is: %lf\n", result);

  return 0;
}
