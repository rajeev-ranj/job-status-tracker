import streamlit as st
import dashboard
import service_management

st.sidebar.title("Navigation")
selection = st.sidebar.radio("Go to", ["Dashboard", "Service Management"])

if selection == "Dashboard":
    dashboard.show_dashboard()
elif selection == "Service Management":
    service_management.show_service_management()
