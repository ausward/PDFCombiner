import tkinter as tk
from tkinter import filedialog, messagebox
import subprocess

# Create the Tkinter root window
root = tk.Tk()
frm = tk.Frame(root)
frm.grid()
tk.Button(frm, text="If there are any issues or question email award@sterlingm.net",).grid(row=0, column=0)




# Ask for a folder using the file dialog
folder_path = filedialog.askdirectory()


# Check if a folder was selected
if folder_path:
    # Start the executable with the folder path as the first argument
    try:
        result = subprocess.run(["./back", folder_path])
    except Exception as e:
        result = subprocess.run(["./back.exe", folder_path])
    if result.returncode != 0:
        # Open tkinter and display an error message
        tk.messagebox.showinfo("Error", "This Folder Contains an encrypted PDF or no PDFs")

root.destroy()  # Close the Tkinter window