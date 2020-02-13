/**
 * Given an array of strictly the characters 'R', 'G', and 'B', 
 * segregate the values of the array so that all the Rs come first, 
 * the Gs come second, and the Bs come last. You can only swap elements of the array. 
 * Do this in linear time and in-place.
 * For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'], 
 * it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].
 */

class SeperateRGB{
    public static void main(String[] args){
        char[] rgb = {'G', 'B', 'R', 'R', 'B', 'R','G'};
        for(char c : rgb){
            System.out.print(c);
        }
        System.out.println();
        int r = 0;
        int b = rgb.length - 1;
        int g = 0;
        char tmp;
        while( g < b){
            if(rgb[g] == 'R'){
                tmp = rgb[r];
                rgb[r] = rgb[g];
                rgb[g] = tmp;
                r++;
            } else if(rgb[g] == 'B') {
                tmp = rgb[b];
                rgb[b] = rgb[g];
                rgb[g] = tmp;
                b--;
            } else {
                g++;
            }
            for(char c : rgb){
                System.out.print(c);
            }
            System.out.println();
        }
        for(char c : rgb){
            System.out.print(c);
        }
        System.out.println();
    }
}