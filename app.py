import streamlit as st

st.title("Welcome to Demo Docker Deployment")

st.header("Simple Calculator")

a= st.number_input("Enter First Number")
b= st.number_input("Enter Second Number")

if st.button:
    c=a+b
    d=a-b
    e=a*b
    st.write("Result of Addition is : ", c)
    st.write("Result of Subtraction is : ", d)
    st.write("Result of Multiplication is : ", e)

d
