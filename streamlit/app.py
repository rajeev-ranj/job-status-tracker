import streamlit as st
import dashboard
import service_management as service

def main():
    st.title("Job Status Tracker")

    # Sidebar for navigation
    menu = st.sidebar.selectbox("Menu", ["Dashboard", "Services", "Jobs", "Users"])

    if menu == "Dashboard":
        dashboard.display_dashboard()
    elif menu == "Services":
        service.display_services()
    elif menu == "Jobs":
        # You can call the corresponding function from the jobs module
        pass
    elif menu == "Users":
        # You can call the corresponding function for managing users
        pass

    st.sidebar.info("Job Status Tracker v1.0")

if __name__ == "__main__":
    main()
