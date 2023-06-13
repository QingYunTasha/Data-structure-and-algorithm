#include <iostream>
#include <vector>
using namespace std;

int main(){
    /*
    if comparing with target, there has three cases:
    1. some true && some false
    2. all true
    3. all false

    if find target, there has two cases:
    1. has target
    2. no target
    */
}

/*
436. Find Right Interval
744. Find Smallest Letter Greater Than Target
1351. Count Negative Numbers in a Sorted Matrix
*/
int findMinIndexGreaterThanTarget(vector<int>& nums, int target) { 
    int l = 0, r = nums.size();

    while (l < r){
        int m = (l + r) >> 1;
        if (nums[m] <= target){
            l = m + 1;
        }else{
            r = m;
        }
    }

    return (l < nums.size()) ? l : -1;
}

/*
1802. Maximum Value at a Given Index in a Bounded Array
*/
int findMaxValLessThanTarget(vector<int>& nums, int target){
    int l = 0, r = nums.size()-1;

    while (l <= r){
        int m = (l + r) >> 1;
        if (nums[m] < target){
            l = m + 1;
        }else{
            r = m - 1;
        }
    }
    return r;
}

/*
704. Binary Search
35. Search Insert Position
*/
int findTarget(vector<int>& nums, int target) {
    int l = 0, r = nums.size()-1;
    while (l <= r){
        int m = (l + r) >> 1;
        if (nums[m] == target){
            return m;
        }else if (nums[m] < target){
            l = m + 1;
        }else{
            r = m - 1;
        }
    }

    return -1;
}

/*
34. Find First and Last Position of Element in Sorted Array
*/
int findLeftmostTarget(vector<int>& nums, int target) {
    int l = 0, r = nums.size()-1;
    while (l <= r) {
        int mid = (l+r) >> 1;
        if (nums[mid] < target)
            l = mid+1;
        else
            r = mid-1;
    }
    return l;
}

/*
34. Find First and Last Position of Element in Sorted Array
*/
int findRightmostTarget(vector<int>& nums, int target) {
    target++;
    int l = 0, r = nums.size()-1;
    while (l <= r) {
        int mid = (l+r) >> 1;
        if (nums[mid] < target)
            l = mid+1;
        else
            r = mid-1;
    }
    return l-1;
}

/*
981. Time Based Key-Value Store
*/
int findMaxIndexLessOrEqualToTarget(vector<int>& nums, int target){
    int l = 0, r = nums.size()-1;

    while (l <= r){
        int m = (l + r) >> 1;
        if (nums[m] <= target){
            l = m + 1;
        }else{
            r = m - 1;
        }
    }
    return r;
}


/*
33. Search in Rotated Sorted Array
*/
int findLeftMinInRotatedArray(vector<int>& nums){
    int left = 0, right = nums.size() - 1;

    while (left < right) {
        int mid = (left + right) >> 1;

        if (nums[mid] > nums[right]) {
            // The original 0 index is on the right side
            left = mid + 1;
        } else if (nums[mid] < nums[right]){
            // The original 0 index is on the left side, or mid itself could be the 0 index
            right = mid;
        }else{
            right--;
        }
    }

    // At this point, left and right will be pointing to the original 0 index
    return left;
}

pair<int, int> findInterval(vector<int>& nums, int target){
    int l = 0, r = nums.size()-1;
    while (l <= r){
            int m = l + (r-l) / 2;
            if (nums[m] > target){
                r = m - 1;
            }else{
                l = m + 1;
            }
        }
        
    return make_pair(r, l);
}