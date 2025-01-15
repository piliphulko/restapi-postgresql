package serviceacctauth

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateAccount(ctx context.Context, in *basic.CustomerNew) (*emptypb.Empty, error) {
	// CHECK LENGTH PASSWORT
	if length := len(in.CustomerAut.PasswordCustomer); length < 8 || length > 64 {
		return &emptypb.Empty{}, apierror.ErrPassLen
	}
	// PASSWORT ENCODING
	pa, err := argon2.CreareArgon2Record([]byte(in.CustomerAut.PasswordCustomer))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	encodedPass, err := pa.Bytes()
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// DATABASE ENTY
	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	var id_customer int
	err = tx.QueryRow(ctx, `
		INSERT INTO table_customer(login_customer, passwort_customer)
		VALUES ($1, $2)
		RETURNING id_customer;`, in.CustomerAut.LoginCustomer, encodedPass).Scan(&id_customer)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO table_customer_info (id_customer, delivery_location_country, delivery_location_city)
		VALUES ($1, $2::enum_country, $3);
	`, id_customer, in.CustomerInfo.CustomerCountry, in.CustomerInfo.CustomerCity)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO table_customer_wallet (id_customer) VALUES ($1);`, id_customer)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// OK
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) CreateAccountWarehouse(ctx context.Context, in *basic.WarehouseNew) (*emptypb.Empty, error) {

	// PASSWORT ENCODING
	pa, err := argon2.CreareArgon2Record([]byte(in.WarehouseAut.PasswordWarehouse))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	encodedPass, err := pa.Bytes()
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// DATABASE ENTY
	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	var id_warehouse int
	err = tx.QueryRow(ctx, `
	INSERT INTO table_warehouse(login_warehouse, passwort_warehouse)
	VALUES ($1, $2)
	RETURNING id_warehouse;`, in.WarehouseAut.LoginWarehouse, encodedPass).Scan(&id_warehouse)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
	INSERT INTO table_warehouse_wallet (id_warehouse) VALUES ($1);`, id_warehouse)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
	INSERT INTO table_warehouse_info (id_warehouse, name_warehouse, info_warehouse, country, city)
	VALUES ($1, $2, $3, $4::enum_country, $5);`,
		id_warehouse, in.WarehouseInfo.WarehouseName, in.WarehouseInfo.WarehouseNote,
		in.WarehouseInfo.WarehouseCountry, in.WarehouseInfo.WarehouseCity)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
	INSERT INTO table_warehouse_commission (id_warehouse, commission_percentage) VALUES ($1, $2::domain_percentage);
`, id_warehouse, in.WarehouseInfo.WarehouseCommission)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// OK
	return nil, status.New(codes.OK, "").Err()
}

func (s *server) CreateAccountVendor(ctx context.Context, in *basic.VendorNew) (*emptypb.Empty, error) {
	// PASSWORT ENCODING
	pa, err := argon2.CreareArgon2Record([]byte(in.VendorAut.PasswordVendor))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	encodedPass, err := pa.Bytes()
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// DATABASE ENTY
	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	var id_vendor int
	err = tx.QueryRow(ctx, `
		INSERT INTO table_vendor(login_vendor, passwort_vendor)
		VALUES ($1, $2)
		RETURNING id_vendor;`, in.VendorAut.LoginVendor, encodedPass).Scan(&id_vendor)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO table_vendor_info (id_vendor, name_vendor)
		VALUES ($1, $2);`, id_vendor, in.VendorInfo.VendorName)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO table_vendor_wallet (id_vendor) VALUES ($1);
	`, id_vendor)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// OK
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()

}
