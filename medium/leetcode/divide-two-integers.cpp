int divide(int dividend, int divisor)
{
    double lna = log(fabs(dividend));
    double lnb = log(fabs(divisor));

    long result = exp(lna-lnb);
    result = (!(dividend<0 ^ divisor<0)) ? result : result*-1;
    if (result>INT_MAX)
        return INT_MAX;
    return result;
}
