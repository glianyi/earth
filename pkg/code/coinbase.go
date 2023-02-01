package code

/**
class Solution {
    public int coinChange(int[] coins, int amount) {
        完全背包之求凑满背包的最少物品数:
        容量->amount; 物品重量->coins[i]; 价值->1; 优化目标->求恰好凑满背包的物品最小价值
        1.状态定义:dp[j]为凑满容量为j的背包所需最少物品数
        2.状态转移:考虑coins[i]
            2.1 当j<coins[i]时:装不下,继承上一个dp[j]的值
            2.2 当j>=coins[i]时:可以装得下,可以选择装或者不装中价值小的(物品数小的)进行转移
                即:dp[j]=min(dp[j],dp[j-coins[i]+1])
        3.初始化:容量为0,最少要装0个就可以装满->dp[0]=0,看转移方程,其他的要初始化为Integer.MAX_VALUE
        4.遍历顺序:这里求最少的物品数,因此排列与组合均可,这里先物品后容量,物品顺序无所谓,容量必须正序(完全背包)
        5.返回形式:返回如果dp[amount]有转移直接返回,如果没有转移返回-1
		 int[] dp = new int[amount + 1];
		 Arrays.fill(dp, Integer.MAX_VALUE);
		 dp[0] = 0;
		 for (int coin : coins) {
			 // 正序,j从coin开始即可
			 for (int j = coin; j <= amount; j++) {
				 // 前面dp值在有计算过的基础上才能转移
				 if(dp[j - coin] != Integer.MAX_VALUE) dp[j] = Math.min(dp[j], dp[j - coin] + 1);
			 }
		 }
		 // 有转移直接返回,没有转移说明凑不成返回-1
		 return dp[amount] == Integer.MAX_VALUE ? -1 : dp[amount];
	 }
 }
*/
