
public class _1_find_var_in_matrix {
    public static void main(String[] args) {
        int[][] matrix = {
                         		{1, 2, 8, 9},
                         		{2, 4, 9, 12},
                         		{4, 7, 10, 13},
                         		{6, 8, 11, 15}
                         	};

        int target = 7;


    }

    public int[] findNuminMatrix(int[][] matrix, int target) {
        int[] result
        int rows = matrix.length;
        int cols = matrix[0].length;

        int x, y = 0, cols-1;

        while (x < rows && y >= 0) {
            int num = matrix[x][y];

            if (target < num) {

            } else if (target > num) {
            } else {
                // equal
            }
        }
    }
}