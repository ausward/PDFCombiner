import tkinter as tk
from tkinter import filedialog, messagebox
import subprocess


def RotatePDF():
    # ask for a single file using the file dialog
    file_path = filedialog.askopenfilename()
    # Check if a file was selected
    if file_path:
        # Start the executable with the file path as the first argument
        try:
            result = subprocess.run(["./back","-R", file_path])
        except Exception as e:
            result = subprocess.run(["./back.exe","-R", file_path])
        if result.returncode != 0:
            # Open tkinter and display an error message
            tk.messagebox.showinfo("Error", "This File Contains an encrypted PDF or no PDFs")


def CombinePDF():
    # Ask for a folder using the file dialog
    folder_path = filedialog.askdirectory()
    # Check if a folder was selected
    if folder_path:
        result = subprocess.CompletedProcess.returncode = 9
        # Start the executable with the folder path as the first argument
        try:
            result = subprocess.run(["./back", folder_path])
        except Exception as e:
            try:
                result = subprocess.run(["./back.exe", folder_path])
            except Exception as e:
                pass
        if result.returncode != 0:
            # Open tkinter and display an error message
            tk.messagebox.showinfo("Error", "This Folder Contains an encrypted PDF or no PDFs")

def RotateAndCombine():
    folder_path = filedialog.askdirectory()
    # Check if a folder was selected
    if folder_path:
        result = subprocess.CompletedProcess.returncode = 9
        # Start the executable with the folder path as the first argument
        try:
            result = subprocess.run(["./back", "-CR", folder_path])
        except Exception as e:
            try:
                result = subprocess.run(["./back.exe", "-CR", folder_path])
            except Exception as e:
                pass
        if result.returncode != 0:
            # Open tkinter and display an error message
            tk.messagebox.showinfo("Error", "This Folder Contains an encrypted PDF or no PDFs")




# Create the Tkinter root window
root = tk.Tk()
frm = tk.Frame(root)
frm.grid()
tk.Button(frm, text="If there are any issues or question email award@sterlingm.net").grid(row=0, column=0)
tk.Button(frm, text="Click here to combine PDFs", command=CombinePDF).grid(row=1, column=0)
tk.Button(frm, text="Click here to rotate PDF", command=RotatePDF).grid(row=2, column=0)
tk.Button(frm, text="Click here to rotate and combine PDFs", command=RotateAndCombine).grid(row=3, column=0)
tk.Button(frm, text="Click here to exit", command=root.destroy).grid(row=4, column=0)


tk.mainloop()  # Start the Tkinter event loop



root.destroy()  # Close the Tkinter window
