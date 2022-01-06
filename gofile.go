package main

import (
  "archive/tar"
  "bytes"
  "io"
  "os"
  "path/filepath"
)

func UnpackArchive(archive []byte, rootPath string) error {
  t := tar.NewReader(bytes.NewReader(archive))

  for {
    hdr, err := t.Next()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      return err
    }
    if hdr == nil {
      continue
    }
    targetPath := filepath.Join(rootPath, hdr.Name)
    switch hdr.Typeflag {
    case tar.TypeDir:
      if _, err := os.Stat(targetPath); err != nil {
        if err := os.MkdirAll(targetPath, os.FileMode(hdr.Mode)); err != nil {
          return err
        }
      }
    case tar.TypeReg:
      f, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
      if err != nil {
        return err
      }
      if _, err := io.Copy(f, t); err != nil {
        return err
      }
      // Not sure why opening and writing would work but closing would fail... but error handling is error handling!
      // Not sure why opening and writing would work but closing would fail... but error handling is error handling!
      if err := f.Close(); err != nil {
        return err
      }
    }
  }
}
