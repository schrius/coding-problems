import java.util.Stack;

/*
 * Given a string of round, curly, and square open and closing brackets, return whether the brackets are balanced (well-formed).
 * For example, given the string "([])[]({})", you should return true.
 * Given the string "([)]" or "((()", you should return false.
 */

 public class ValidParathese{
    public static void main(String[] args) {
        String expression1 = "([])[]({})";
        String expression2 = "([)]";
        String expression3 = "((()";

        if(validParathese(expression1)){
            System.out.println("Valid");
        } else {
            System.out.println("Invalid");
        }
        if(validParathese(expression2)){
            System.out.println("Valid");
        } else {
            System.out.println("Invalid");
        }
        if(validParathese(expression3)){
            System.out.println("Valid");
        } else {
            System.out.println("Invalid");
        }
    }

    static boolean validParathese(String expression){
        Stack<Character> stack = new Stack<>();
        if(expression.length() <= 1 || expression.length()/2 == 1){
            System.out.println("Invalid length");
            return false;
        }
        
        for(int i = 0 ; i < expression.length(); i++){
            if(expression.charAt(i) == '(' || expression.charAt(i) == '{'  || expression.charAt(i) == '[' )
                stack.push(expression.charAt(i));
            else if(!stack.isEmpty()){
                if(stack.pop().equals(expression.charAt(i))){
                    return false;
                }
            }
            else return false;
        }
        return true;
    }
 }