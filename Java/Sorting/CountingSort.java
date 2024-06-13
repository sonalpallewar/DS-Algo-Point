package Java.Sorting;

import java.util.*;
import java.io.*;

class CountingSort {
    
    public static void main(String args[])throws IOException{
      BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
      int n = Integer.parseInt(reader.readLine());         //input array size
      int arr[] = new int[n];                              //input array
      StringTokenizer tk = new StringTokenizer(reader.readLine());
      for(int i=0;i<n;i++){
        arr[i]=Integer.parseInt(tk.nextToken());
      }
      System.out.println("Array before sorting-");
      for(int i=0;i<n;i++){
          System.out.print(arr[i]+" ");  
      }
      System.out.println("");
      sort(arr,n);                 //utility function for counting sort
      System.out.println("Array after sorting-");
      for(int i=0;i<n;i++){
          System.out.print(arr[i]+" ");  
      }
    }
}
/*Counting sort 

-:Input/Output Format:-

Sample Input-
5
7 2 5 3 8
Expected Output-
Array before sorting- 
7 2 5 3 8
Array after sorting- 
2 3 5 7 8

-:Complexities:-

Time Complexity is:O(n+max)
Space Complexity is:O(n+max)
                    where max is the largest number in the input array
                    and n is the size of input array*/ 